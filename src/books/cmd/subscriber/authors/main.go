package main

import (
	"context"
	"go-boilerplate/shared/event"
	"go-boilerplate/shared/log"
	"go-boilerplate/src/books/domain/authors"
	"go-boilerplate/src/books/services"
	"go-boilerplate/src/books/services/command"

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

	app := services.NewBookService()

	subscriber := event.InitSubscriber(event.SubscriberOptions{Topic: "authors"})

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
}

func handleDeleteAuthor(c context.Context, params HandlerParams, event event.Event) error {
	var data authors.DeleteAuthorEvent
	err := event.TransformTo(&data)
	if err != nil {
		params.logger.Errorf("Failed to transform event data: %v", tracerr.Wrap(err))
	}

	err = params.app.Commands.DeleteBookByAuthor.Execute(c, command.DeleteBookByAuthorParams{
		AuthorID: data.ID,
	})
	if err != nil {
		params.logger.Errorf("Failed to delete books by author: %v", tracerr.Wrap(err))
		return err
	}

	return nil
}
