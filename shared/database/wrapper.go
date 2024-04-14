package database

import (
	"context"
	"database/sql"
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
	res := t.tx.QueryRowContext(ctx, query, args...)

	return res
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
	tx, err := db.DB.BeginTx(c, options)
	if err != nil {
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

func (db DB) ExecContext(c context.Context, query string, args ...interface{}) (sql.Result, error) {
	return TxFromContext(c, db.DB).ExecContext(c, query, args...)
}

func (db DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db DB) QueryContext(c context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return TxFromContext(c, db.DB).QueryContext(c, query, args...)
}

func (db DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}

func (db DB) QueryRowContext(c context.Context, query string, args ...interface{}) *sql.Row {
	return TxFromContext(c, db.DB).QueryRowContext(c, query, args...)
}
