package pagination

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
)

type Metadata struct {
	Next *uuid.UUID
	Prev *uuid.UUID
}

type Collection[Entity any] struct {
	Data     []Entity
	Metadata Metadata
}

type Config[T any] interface {
	Paginate(conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[T], error)
}

func NewPaginate[Entity any](strategy Config[Entity], conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[Entity], error) {
	return strategy.Paginate(conn, fn)
}
