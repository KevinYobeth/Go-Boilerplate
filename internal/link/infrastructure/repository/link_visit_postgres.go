package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
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
		Values(dto.ID, dto.LinkID, dto.Total, nil).
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

func (r *PostgresLinkRepo) GetNewVisitsCount(c context.Context) ([]link.NewVisitCountModel, error) {
	snapshots := psql.Select("link_id", "total", "last_snapshot_at").
		From("link_visit_snapshots")

	query, args, err := psql.
		Select("v.link_id", "COUNT(*) AS new_visits", "MAX(v.visited_at) AS latest_visit").
		FromSelect(snapshots, "s").
		Join("link_visits v ON v.link_id = s.link_id").
		Where(sq.Or{
			sq.Expr("last_snapshot_at IS NULL"),
			sq.Expr("visited_at > last_snapshot_at"),
		}).
		GroupBy("v.link_id").
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	rows, err := r.db.QueryContext(c, query, args...)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	defer rows.Close()

	var newVisitCounts []link.NewVisitCountModel
	for rows.Next() {
		var newVisitCount link.NewVisitCountModel
		err = rows.Scan(&newVisitCount.LinkID, &newVisitCount.NewVisits, &newVisitCount.LatestVisit)
		if err != nil {
			return nil, tracerr.Wrap(err)
		}

		newVisitCounts = append(newVisitCounts, newVisitCount)
	}

	return newVisitCounts, nil
}

func (r *PostgresLinkRepo) UpdateLinkVisitSnapshot(c context.Context, dto []link.UpdateLinkVisitSnapshotDTO) error {
	if len(dto) == 0 {
		return nil
	}

	caseStmt := sq.Case("link_id")
	linkIDs := make([]string, 0, len(dto))

	for _, v := range dto {
		caseStmt = caseStmt.When(fmt.Sprintf("'%s'", v.LinkID.String()), sq.Expr("total + ?", v.NewVisits))
		linkIDs = append(linkIDs, v.LinkID.String())
	}

	now := time.Now()

	query, args, err := psql.Update("link_visit_snapshots").
		Set("total", caseStmt).
		Set("last_snapshot_at", now).
		Where(sq.Eq{"link_id": linkIDs}).
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
