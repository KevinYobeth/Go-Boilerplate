package repository

import (
	"context"
	"database/sql"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/books/domain/books"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresBooksRepo struct {
	db database.PostgresDB
}

func NewBooksPostgresRepository(db database.PostgresDB) Repository {
	return &PostgresBooksRepo{db}
}

func (r PostgresBooksRepo) GetBooks(c context.Context, request books.GetBooksDto) ([]books.Book, error) {
	query, args, err := psql.Select("id", "title").
		From("books").
		Where(sq.Eq{"deleted_at": nil}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	rows, err := r.db.QueryContext(c, query, args...)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	defer rows.Close()

	var booksResult []books.Book
	for rows.Next() {
		var book books.Book
		rows.Scan(&book.ID, &book.Title)

		booksResult = append(booksResult, book)
	}

	return booksResult, nil
}

func (r PostgresBooksRepo) GetBook(c context.Context, id uuid.UUID) (*books.Book, error) {
	query, args, err := psql.Select("id", "title").
		From("books").
		Where(sq.Eq{"id": id, "deleted_at": nil}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	var book books.Book
	err = r.db.QueryRowContext(c, query, args...).
		Scan(&book.ID, &book.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	return &book, nil
}

func (r PostgresBooksRepo) CreateBook(c context.Context, request books.CreateBookDto) error {
	now := time.Now().UTC()

	query, args, err := psql.Insert("books").
		SetMap(map[string]interface{}{
			"id":         request.ID,
			"title":      request.Title,
			"created_at": now,
		}).
		ToSql()
	if err != nil {
		return tracerr.Wrap(err)
	}

	_, err = r.db.ExecContext(c, query, args...)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (r PostgresBooksRepo) UpdateBook(c context.Context, request books.UpdateBookDto) error {
	now := time.Now().UTC()

	query, args, err := psql.Update("books").
		SetMap(map[string]interface{}{
			"title":      request.Title,
			"updated_at": now,
		}).
		Where(sq.Eq{"id": request.ID}).
		ToSql()
	if err != nil {
		return tracerr.Wrap(err)
	}

	_, err = r.db.ExecContext(c, query, args...)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (r PostgresBooksRepo) DeleteBook(c context.Context, id uuid.UUID) error {
	now := time.Now().UTC()

	query, args, err := psql.Update("books").
		Where(sq.Eq{"id": id}).
		SetMap(map[string]interface{}{
			"deleted_at": now,
		}).ToSql()

	if err != nil {
		return tracerr.Wrap(err)
	}

	_, err = r.db.ExecContext(c, query, args...)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (r PostgresBooksRepo) CreateAuthorBook(c context.Context, request books.CreateAuthorBookDto) error {
	query, args, err := psql.Insert("author_book").
		Columns("book_id", "author_id").
		Values(request.BookID, request.AuthorID).
		ToSql()
	if err != nil {
		return tracerr.Wrap(err)
	}

	_, err = r.db.ExecContext(c, query, args...)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
