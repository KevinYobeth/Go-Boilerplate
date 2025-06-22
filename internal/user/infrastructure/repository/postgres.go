package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"
	"github.com/ztrue/tracerr"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresUserRepo struct {
	db database.PostgresDB
}

func NewUserPostgresRepository(db database.PostgresDB) Repository {
	return &PostgresUserRepo{db}
}

func (r *PostgresUserRepo) GetUser(c context.Context, id uuid.UUID) (*user.User, error) {
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

func (r *PostgresUserRepo) GetUserByEmail(c context.Context, email string) (*user.User, error) {
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
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	return user, nil
}
