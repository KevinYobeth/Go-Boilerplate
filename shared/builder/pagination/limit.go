package pagination

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/scan/v2"
	"github.com/ztrue/tracerr"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type LimitPaginationRequest[Entity EntityUniqueColumn] struct {
	Page    int
	Limit   int
	OrderBy string
}

func NewLimitPagination[Entity EntityUniqueColumn](request LimitPaginationRequest[Entity]) Config[Entity] {
	var page int = 1
	var limit int = 5
	var order string = "created_at"

	if request.Page != 0 {
		page = request.Page
	}

	if request.Limit != 0 {
		limit = request.Limit
	}

	if request.OrderBy != "" {
		order = request.OrderBy
	}

	return &LimitPaginationRequest[Entity]{
		Page:    page,
		Limit:   limit,
		OrderBy: order,
	}
}

func (l *LimitPaginationRequest[Entity]) Paginate(ctx context.Context, conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[Entity], error) {
	var collection Collection[Entity]
	var base = fn(conn)

	err := ValidateLimitPaginationParams(l.Page, l.Limit)
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	query, args, err := psql.
		Select("COUNT(1)").
		FromSelect(base, "sub").
		ToSql()
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	var total *uint64
	row := conn.QueryRowContext(ctx, query, args...)
	err = row.Scan(&total)
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	totalInt := int(*total)
	collection.Metadata.Total = &totalInt
	collection.Metadata.Limit = &l.Limit
	collection.Metadata.Page = &l.Page

	result := base.
		Limit(uint64(l.Limit)).
		Offset((uint64(l.Page-1) * uint64(l.Limit))).
		OrderBy(l.OrderBy)

	query, args, err = result.ToSql()
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	var items []Entity
	err = scan.Rows(&items, rows)
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	collection.Data = items

	return collection, nil
}
