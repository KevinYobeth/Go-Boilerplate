package pagination

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/scan/v2"
	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/database"
	"github.com/ztrue/tracerr"
)

type Uniqueable interface {
	GetID() any
}

type CursorPaginationRequest[Entity Uniqueable] struct {
	Limit uint64
	Order string
	Next  uuid.UUID
	Prev  uuid.UUID
}

func NewCursorPagination[Entity Uniqueable](request CursorPaginationRequest[Entity]) Config[Entity] {
	var limit uint64 = 5
	var order string = "created_at"

	if request.Limit != 0 {
		limit = request.Limit
	}

	if request.Order != "" {
		order = request.Order
	}

	return &CursorPaginationRequest[Entity]{
		Limit: limit,
		Order: order,
		Next:  request.Next,
		Prev:  request.Prev,
	}
}

func (l *CursorPaginationRequest[Entity]) Paginate(ctx context.Context, conn database.PostgresDB, fn func(conn database.PostgresDB) sq.SelectBuilder) (Collection[Entity], error) {
	var collection Collection[Entity]
	var base = fn(conn).
		Limit(l.Limit + 1).
		OrderBy(l.Order)

	if l.Next != uuid.Nil && l.Prev != uuid.Nil {
		// TODO: Check if both Next and Prev are set
	}

	if l.Prev != uuid.Nil {
		result, err := l.getCurrentCursor(conn, fn(conn), l.Prev.String())
		if err != nil {
			return collection, tracerr.Wrap(err)
		}

		base = base.Where("(created_at, id) < (?, ?)", result["created_at"], l.Prev)
	}

	if l.Next != uuid.Nil {
		result, err := l.getCurrentCursor(conn, fn(conn), l.Next.String())
		if err != nil {
			return collection, tracerr.Wrap(err)
		}

		base = base.Where("(created_at, id) > (?, ?)", result["created_at"], l.Next)
	}

	query, args, err := base.ToSql()
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return collection, tracerr.Wrap(err)
	}
	defer rows.Close()

	var items []Entity
	err = scan.Rows(&items, rows)
	if err != nil {
		return collection, tracerr.Wrap(err)
	}

	collection.Data = items[:len(items)-1]

	if l.Next != uuid.Nil {
		collection.Metadata.Next = items[len(items)-2].GetID().(*uuid.UUID)

		return collection, nil
	}

	if l.Prev != uuid.Nil {
	}

	collection.Metadata.Next = items[len(items)-2].GetID().(*uuid.UUID)

	return collection, nil
}

func (l *CursorPaginationRequest[Entity]) scanToRows(rows *sql.Rows) ([]map[string]any, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	result := make([]map[string]any, 0)

	for rows.Next() {
		columns := make([]any, len(cols))
		pointers := make([]any, len(cols))

		for i := range columns {
			pointers[i] = &columns[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		model := make(map[string]any)

		for i, it := range cols {
			model[it] = *pointers[i].(*any)
		}

		result = append(result, model)
	}

	if err := rows.Err(); err != nil {
		return nil, tracerr.Wrap(err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (l *CursorPaginationRequest[Entity]) getCurrentCursor(conn database.PostgresDB, query sq.SelectBuilder, cursor string) (map[string]any, error) {
	qry, args, err := query.
		Where(sq.Eq{"id": cursor}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	rows, err := conn.QueryContext(context.Background(), qry, args...)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	result, err := l.scanToRows(rows)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return result[0], nil
}
