package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/utils"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresAuthenticationRepo struct {
	db database.PostgresDB
}

func NewAuthenticationPostgresRepository(db database.PostgresDB) Repository {
	return &PostgresAuthenticationRepo{db}
}

func (r *PostgresAuthenticationRepo) Register(c context.Context, dto *user.UserDto) error {
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

func (r *PostgresAuthenticationRepo) GetUser(c context.Context, id uuid.UUID) (*user.User, error) {
	fields := utils.SelectWithAuditTrail("id", "first_name", "last_name", "email", "password")
	query, args, err := psql.Select(fields...).
		From("users").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	row := r.db.QueryRowContext(c, query, args...)

	user := &user.User{}
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return user, nil
}

func (r *PostgresAuthenticationRepo) GetUserByEmail(c context.Context, email string) (*user.User, error) {
	fields := utils.SelectWithAuditTrail("id", "first_name", "last_name", "email", "password")
	query, args, err := psql.Select(fields...).
		From("users").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	row := r.db.QueryRowContext(c, query, args...)

	user := &user.User{}
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy, &user.UpdatedBy)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return user, nil
}
