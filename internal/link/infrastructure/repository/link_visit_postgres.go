package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/ztrue/tracerr"
)

func (r *PostgresLinkRepo) CreateLinkVisit(c context.Context, dto *link.LinkVisitEventDTO) error {
	query, args, err := psql.Insert("link_visits").
		Columns("id", "link_id", "slug", "ip_address", "user_agent").
		Values(dto.ID, dto.LinkID, dto.Slug, dto.IPAddress, dto.UserAgent).
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

func (r *PostgresLinkRepo) CreateLinkVisitSnapshot(c context.Context, dto *link.LinkVisitSnapshotDTO) error {
	query, args, err := psql.Insert("link_visit_snapshots").
		Columns("id", "link_id", "total", "last_snapshot_at").
		Values(dto.ID, dto.LinkID, dto.Total, dto.LastSnapshotAt).
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
