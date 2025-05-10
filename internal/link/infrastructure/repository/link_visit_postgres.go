package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/ztrue/tracerr"
)

func (r *PostgresLinkRepo) CreateLinkVisitEvent(c context.Context, dto *link.LinkVisitEventDTO) error {
	query, args, err := psql.Insert("link_visit_events").
		Columns("id", "link_id", "ip_address", "user_agent", "referer", "country_code", "device_type", "browser").
		Values(dto.ID, dto.LinkID, dto.IPAddress, dto.UserAgent, dto.Referer, dto.CountryCode, dto.DeviceType, dto.Browser).
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
