package database

import (
	"context"
	"database/sql"

	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"

	"go.opentelemetry.io/otel/trace"
)

type DB struct {
	DB *sql.DB
}

type Tx struct {
	tx *sql.Tx
}

func (t *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.tx.Exec(query, args...)
}

func (t *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return t.tx.Query(query, args...)
}

func (t Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRow(query, args...)
}

func (t Tx) QueryContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (res *sql.Rows, err error) {
	return t.tx.QueryContext(ctx, query, args...)
}

func (t Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t Tx) ExecContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (res sql.Result, err error) {
	return t.tx.ExecContext(ctx, query, args...)
}

func (t Tx) Rollback() error {
	return t.tx.Rollback()
}

func (t Tx) Commit() error {
	return t.tx.Commit()
}

func NewDB(db *sql.DB) DB {
	return DB{DB: db}
}

// BeginTx implements PostgresDB.
func (db DB) BeginTx(c context.Context, options *sql.TxOptions) (Transaction, error) {
	ctx, span := telemetry.NewDatabaseSpan(c, "BeginTx")
	defer span.End()

	tx, err := db.DB.BeginTx(ctx, options)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		return nil, err
	}

	return &Tx{
		tx: tx,
	}, nil
}

func (db DB) Ping() error {
	return db.DB.Ping()
}

func (db DB) PingContext(c context.Context) error {
	return db.DB.PingContext(c)
}

func (db DB) Close() error {
	return db.DB.Close()
}

func (db DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db DB) ExecContext(c context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	ctx, span := telemetry.NewDatabaseSpan(c, query)
	defer span.End()

	res, err = TxFromContext(ctx, db.DB).ExecContext(ctx, query, args...)

	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
	}

	return res, err
}

func (db DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db DB) QueryContext(c context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	ctx, span := telemetry.NewDatabaseSpan(c, query)
	defer span.End()

	rows, err = TxFromContext(ctx, db.DB).QueryContext(ctx, query, args...)

	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
	}

	return rows, err
}

func (db DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}

func (db DB) QueryRowContext(c context.Context, query string, args ...interface{}) *sql.Row {
	ctx, span := telemetry.NewDatabaseSpan(c, query)
	defer span.End()

	return TxFromContext(ctx, db.DB).QueryRowContext(ctx, query, args...)
}
