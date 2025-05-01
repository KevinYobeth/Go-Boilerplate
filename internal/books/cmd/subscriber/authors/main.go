package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	authorIntraprocess "github.com/kevinyobeth/go-boilerplate/internal/authors/presentation/intraprocess"
	authorsService "github.com/kevinyobeth/go-boilerplate/internal/authors/services"
	"github.com/kevinyobeth/go-boilerplate/internal/books/domain/authors"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/intraprocess"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/command"
	"github.com/kevinyobeth/go-boilerplate/shared/event"
	"github.com/kevinyobeth/go-boilerplate/shared/graceroutine"
	"github.com/kevinyobeth/go-boilerplate/shared/log"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type HandlerParams struct {
	logger *zap.SugaredLogger
	app    *services.Application
}

func main() {
	logger := log.InitLogger()
	c := context.Background()

	authorsService := authorsService.NewAuthorService()
	authorIntraprocess := authorIntraprocess.NewAuthorIntraprocessService(authorsService)

	booksAuthorIntraprocess := intraprocess.NewBookAuthorIntraprocessService(authorIntraprocess)

	app := services.NewBookService(booksAuthorIntraprocess)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	subscriber := event.InitSubscriber(event.SubscriberOptions{Topic: "authors"})

	go func() {
		subscriber.Subscribe(c, func(ctx context.Context, e event.Event) error {
			logger.Infof("Received event: %s", e.Event)
			var err error

			switch e.Event {
			case "author.delete":
				err = handleDeleteAuthor(ctx, HandlerParams{logger, &app}, e)
			default:
				logger.Infof("Event %s is not handled", e.Event)
			}

			return tracerr.Wrap(err)
		})
	}()

	<-signals

	logger.Info("Shutting down subscriber...")

	if err := subscriber.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Subscriber Shutdown")
}

func handleDeleteAuthor(c context.Context, params HandlerParams, event event.Event) error {
	var data authors.DeleteAuthorEvent
	err := event.TransformTo(&data)
	if err != nil {
		params.logger.Errorf("Failed to transform event data: %v", tracerr.Wrap(err))
	}

	err = params.app.Commands.DeleteBookByAuthor.Handle(c, command.DeleteBookByAuthorRequest{
		AuthorID: data.ID,
	})
	if err != nil {
		params.logger.Errorf("Failed to delete books by author: %v", tracerr.Wrap(err))
		return err
	}

	return nil
}
