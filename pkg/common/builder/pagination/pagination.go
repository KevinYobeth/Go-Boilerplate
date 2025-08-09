package pagination

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
)

type Metadata struct {
	Next    *uuid.UUID
	Prev    *uuid.UUID
	Total   *uint64
	HasNext bool
	HasPrev bool
}

type Collection[Entity any] struct {
	Data     []Entity
	Metadata Metadata
}

type Config[T any] interface {
	Paginate(ctx context.Context, conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[T], error)
}

func NewPaginate[Entity any](ctx context.Context, strategy Config[Entity], conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[Entity], error) {
	return strategy.Paginate(ctx, conn, fn)
}
