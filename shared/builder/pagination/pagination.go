package pagination

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/shared/database"
)

type Metadata struct {
	Next    *uuid.UUID
	Prev    *uuid.UUID
	Total   *int
	Limit   *int
	Page    *int
	HasNext bool
	HasPrev bool
}

type Collection[Entity any] struct {
	Data     []Entity
	Metadata Metadata
}

type Config[T EntityUniqueColumn] interface {
	Paginate(ctx context.Context, conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[T], error)
}

func NewPaginate[Entity EntityUniqueColumn](ctx context.Context, strategy Config[Entity], conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[Entity], error) {
	return strategy.Paginate(ctx, conn, fn)
}

type EntityUniqueColumn interface {
	UniqueColumn() string
}
