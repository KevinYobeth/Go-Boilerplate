package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
	"github.com/kevinyobeth/go-boilerplate/shared/utils"
	"github.com/ztrue/tracerr"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PostgresLinkRepo struct {
	db database.PostgresDB
}

func NewLinkPostgresRepository(db database.PostgresDB) Repository {
	return &PostgresLinkRepo{db}
}

func (r *PostgresLinkRepo) CreateLink(c context.Context, dto *link.LinkDTO) error {
	query, args, err := psql.Insert("links").
		Columns("id", "slug", "url", "description", "created_by").
		Values(dto.ID, dto.Slug, dto.URL, dto.Description, dto.CreatedBy).
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

func (r *PostgresLinkRepo) GetLinks(c context.Context, userID uuid.UUID) ([]link.Link, error) {
	fields := utils.SelectWithAuditTrail("id", "slug", "url", "description")
	query, args, err := psql.Select(fields...).
		From("links").
		Where(sq.Eq{"created_by": userID, "deleted_at": nil}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	rows, err := r.db.QueryContext(c, query, args...)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	defer rows.Close()

	var linksResult []link.Link
	for rows.Next() {
		var link link.Link
		err = rows.Scan(&link.ID, &link.Slug, &link.URL, &link.Description, &link.CreatedAt, &link.UpdatedAt, &link.CreatedBy, &link.UpdatedBy)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		linksResult = append(linksResult, link)
	}

	return linksResult, nil
}

func (r *PostgresLinkRepo) GetLinkBySlug(c context.Context, slug string) (*link.RedirectLink, error) {
	query, args, err := psql.Select("id", "slug", "url").
		From("links").
		Where(sq.Eq{"slug": slug}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	row := r.db.QueryRowContext(c, query, args...)

	link := &link.RedirectLink{}
	err = row.Scan(&link.ID, &link.Slug, &link.URL)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return link, nil
}
