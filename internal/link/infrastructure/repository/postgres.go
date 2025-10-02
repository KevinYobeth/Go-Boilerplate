package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/shared/builder/pagination"
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

func (r *PostgresLinkRepo) GetLinks(c context.Context, userID uuid.UUID) ([]link.LinkModel, error) {
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

	var linksResult []link.LinkModel
	for rows.Next() {
		var link link.LinkModel
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
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	return link, nil
}

func (r *PostgresLinkRepo) GetLink(c context.Context, id, userID uuid.UUID) (*link.Link, error) {
	fields := utils.SelectWithAuditTrail("id", "slug", "url", "description")
	query, args, err := psql.Select(fields...).
		From("links").
		Where(sq.Eq{"id": id, "created_by": userID, "deleted_at": nil}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	row := r.db.QueryRowContext(c, query, args...)
	link := &link.Link{}
	err = row.Scan(&link.ID, &link.Slug, &link.URL, &link.Description, &link.CreatedAt, &link.UpdatedAt, &link.CreatedBy, &link.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, tracerr.Wrap(err)
	}

	return link, nil
}

func (r *PostgresLinkRepo) DeleteLink(c context.Context, id uuid.UUID) error {
	now := time.Now().UTC()

	query, args, err := psql.Update("links").
		Set("deleted_at", now).
		Where(sq.Eq{"id": id}).
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

func (r *PostgresLinkRepo) UpdateLink(c context.Context, id uuid.UUID, dto *link.LinkDTO) error {
	query, args, err := psql.Update("links").
		Set("slug", dto.Slug).
		Set("url", dto.URL).
		Set("description", dto.Description).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": id}).
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

func (r *PostgresLinkRepo) GetLinksVisitSnapshot(c context.Context, linkIDs []uuid.UUID) ([]link.LinkVisitSnapshot, error) {
	query, args, err := psql.Select("id", "link_id", "total", "last_snapshot_at").
		From("link_visit_snapshots").
		Where(sq.Eq{"link_id": linkIDs}).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	rows, err := r.db.QueryContext(c, query, args...)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	defer rows.Close()

	var linkVisitSnapshotResult []link.LinkVisitSnapshot
	for rows.Next() {
		var link link.LinkVisitSnapshot
		err = rows.Scan(&link.ID, &link.LinkID, &link.Total, &link.LastSnapshotAt)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		linkVisitSnapshotResult = append(linkVisitSnapshotResult, link)
	}

	return linkVisitSnapshotResult, nil
}

func (r *PostgresLinkRepo) GetLinksPaginated(c context.Context, userID uuid.UUID, config pagination.Config[link.LinkModel]) (pagination.Collection[link.LinkModel], error) {
	collection, err := pagination.NewPaginate(c, config, r.db, func(conn database.PostgresDB) sq.SelectBuilder {
		fields := utils.SelectWithAuditTrail("id", "slug", "url", "description")
		query := psql.Select(fields...).
			From("links").
			Where(sq.Eq{"created_by": userID, "deleted_at": nil})

		return query
	})
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	return collection, nil
}
