package repository

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/authors/domain/authors"

	sq "github.com/Masterminds/squirrel"
	"github.com/ztrue/tracerr"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresAuthorsRepo struct {
	db database.PostgresDB
}

func NewAuthorsPostgresRepository(db database.PostgresDB) Repository {
	return &PostgresAuthorsRepo{db}
}

func (r *PostgresAuthorsRepo) GetAuthors(c context.Context, request authors.GetAuthorsDto) ([]authors.Author, error) {
	query, args, err := psql.Select("id", "name").
		From("authors").
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

	var authorsResult []authors.Author
	for rows.Next() {
		var author authors.Author
		rows.Scan(&author.ID, &author.Name)

		authorsResult = append(authorsResult, author)
	}

	return authorsResult, nil
}

func (p *PostgresAuthorsRepo) GetAuthor(c context.Context, id string) (*authors.Author, error) {
	query, args, err := psql.Select("id", "name").
		From("authors").
		Where(sq.Eq{"id": id, "deleted_at": nil}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	row := p.db.QueryRowContext(c, query, args...)
	var author authors.Author
	err = row.Scan(&author.ID, &author.Name)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &author, nil
}

func (p *PostgresAuthorsRepo) CreateAuthor(c context.Context, request authors.CreateAuthorDto) error {
	query, args, err := psql.Insert("authors").
		Columns("id", "name").
		Values(request.ID, request.Name).
		ToSql()
	if err != nil {
		return tracerr.Wrap(err)
	}

	_, err = p.db.ExecContext(c, query, args...)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
