package main

import (
	"context"
	"fmt"
	"go-boilerplate/shared/event"
	"go-boilerplate/shared/log"
	"go-boilerplate/src/books/services"
	"go-boilerplate/src/books/services/command"
	"go-boilerplate/src/books/services/query"

	"go.uber.org/zap"
)

type HandlerParams struct {
	logger *zap.SugaredLogger
	app    *services.Application
}

func main() {
	logger := log.InitLogger()

	subscriber := event.InitSubscriber(event.SubscriberOptions{Topic: "authors"})
	app := services.NewBookService()

	c := context.Background()

	subscriber.Subscribe(c, func(ctx context.Context, e event.Event) error {
		var err error

		switch e.Event {
		case "author.delete":
			err = handleDeleteAuthor(ctx, HandlerParams{logger, &app}, e)
		default:
			logger.Infof("Event %s is not handled", e.Event)
		}

		return err
	})
}

func handleDeleteAuthor(c context.Context, params HandlerParams, event event.Event) error {
	fmt.Println("EVENT DIDALEM", event)
	books, err := params.app.Queries.GetBooks.Execute(c, query.GetBooksParams{})

	if err != nil {
		params.logger.Errorf("Failed to get books: %v", err)
		return err
	}

	for _, book := range books {
		err := params.app.Commands.DeleteBook.Execute(c, command.DeleteBookParams{ID: book.ID})
		if err != nil {
			params.logger.Errorf("Failed to delete book: %v", err)
			return err
		}
	}

	return nil
}
