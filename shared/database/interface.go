package database

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type (
	BeforeFunc func(ctx context.Context, query string, args ...interface{}) context.Context
	AfterFunc  func(ctx context.Context, err error, query string, args ...interface{})
)

type Queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Transaction interface {
	sq.StdSqlCtx
	Rollback() error
	Commit() error
}

type SqlDatabase interface {
	sq.StdSqlCtx
	BeginTx(context.Context, *sql.TxOptions) (Transaction, error)
	Ping() error
	PingContext(ctx context.Context) error
	Close() error
}
