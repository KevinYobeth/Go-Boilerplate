package repository

import (
	"context"
	"go-boilerplate/shared/database"
	"go-boilerplate/src/authentication/domain/user"

	sq "github.com/Masterminds/squirrel"
	"github.com/ztrue/tracerr"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresAuthenticationRepo struct {
	db database.PostgresDB
}

func NewAuthenticationPostgresRepository(db database.PostgresDB) Repository {
	return &PostgresAuthenticationRepo{db}
}

func (r *PostgresAuthenticationRepo) Register(c context.Context, dto *user.RegisterDto) error {
	query, args, err := psql.Insert("users").
		Columns("id", "first_name", "last_name", "email", "password", "created_by").
		Values(dto.ID, dto.FirstName, dto.LastName, dto.Email, dto.Password, dto.ID).
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
