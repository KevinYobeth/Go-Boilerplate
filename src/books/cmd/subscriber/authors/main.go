package main

import (
	"context"
	"go-boilerplate/shared/event"
	"go-boilerplate/shared/log"
	"go-boilerplate/src/books/services"
	"go-boilerplate/src/books/services/command"
	"go-boilerplate/src/books/services/query"
)

func main() {
	logger := log.InitLogger()
	subscriber := event.InitSubscriber("authors")
	app := services.NewBookService()

	c := context.Background()

	subscriber.Subscribe(c, func(ctx context.Context, e event.Event) error {
		books, err := app.Queries.GetBooks.Execute(c, query.GetBooksParams{})

		if err != nil {
			logger.Errorf("Failed to get books: %v", err)
			return err
		}

		for _, book := range books {
			err := app.Commands.DeleteBook.Execute(c, command.DeleteBookParams{ID: book.ID})
			if err != nil {
				logger.Errorf("Failed to delete book: %v", err)
				return err
			}
		}

		return nil
	})
}
