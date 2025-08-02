package pagination

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
)

type LimitPaginationRequest[Entity any] struct {
	Page    uint64
	Limit   uint64
	OrderBy string
}

func NewLimitPagination[Entity any](request LimitPaginationRequest[Entity]) Config[Entity] {
	var page uint64 = 1
	var limit uint64 = 5
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

func (l *LimitPaginationRequest[Entity]) Paginate(conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[Entity], error) {
	// query = query.Offset(l.Page * l.Limit)
	// query = query.Limit(l.Limit)
	// query = query.OrderBy(l.OrderBy)

	// return query

	return Collection[Entity]{}, nil
}
