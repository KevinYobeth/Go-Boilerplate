package pagination

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/scan/v2"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/ztrue/tracerr"
)

type LimitPaginationRequest[Entity any] struct {
	Page    *uint64
	Limit   *uint64
	OrderBy string
}

func NewLimitPagination[Entity any](request LimitPaginationRequest[Entity]) Config[Entity] {
	var page uint64 = 1
	var limit uint64 = 5
	var order string = "created_at"

	if request.Page != nil {
		page = *request.Page
	}

	if request.Limit != nil {
		limit = *request.Limit
	}

	if request.OrderBy != "" {
		order = request.OrderBy
	}

	return &LimitPaginationRequest[Entity]{
		Page:    &page,
		Limit:   &limit,
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

	query, args, err := base.
		RemoveColumns().
		Column("COUNT(id)").
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
	fmt.Println("totalnya", total)

	collection.Metadata.Total = total

	result := base.
		Limit(*l.Limit).
		Offset((*l.Page - 1) * *l.Limit).
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
