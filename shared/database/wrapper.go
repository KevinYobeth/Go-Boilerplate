package database

import (
	"context"
	"database/sql"

	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"

	"go.opentelemetry.io/otel/trace"
)

type DB struct {
	DB     *sql.DB
	before []BeforeFunc
	after  []AfterFunc
}

type Tx struct {
	tx     *sql.Tx
	before []BeforeFunc
	after  []AfterFunc
}

func (t *Tx) Exec(query string, args ...interface{}) (res sql.Result, err error) {
	callerName := utils.GetFnCallerName(1, 2)
	ctx := context.WithValue(
		context.Background(),
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)
	t.doBefore(ctx, query, args...)
	defer t.doAfter(ctx, err, query, args...)

	res, err = t.tx.Exec(query, args...)

	return
}

func (t *Tx) Query(query string, args ...interface{}) (res *sql.Rows, err error) {
	callerName := utils.GetFnCallerName(1, 2)
	ctx := context.WithValue(
		context.Background(),
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)
	t.doBefore(ctx, query, args...)
	defer t.doAfter(ctx, err, query, args...)

	res, err = t.tx.Query(query, args...)

	return
}

func (t Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	callerName := utils.GetFnCallerName(1, 2)
	ctx := context.WithValue(
		context.Background(),
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)
	t.doBefore(ctx, query, args...)

	res := t.tx.QueryRow(query, args...)
	err := res.Err()

	defer t.doAfter(ctx, err, query, args...)

	return res
}

func (t Tx) QueryContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (res *sql.Rows, err error) {
	callerName := utils.GetFnCallerName(1, 2)
	ctx = context.WithValue(
		ctx,
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)
	t.doBefore(ctx, query, args...)
	defer t.doAfter(ctx, err, query, args...)

	res, err = t.tx.QueryContext(ctx, query, args...)

	return
}

func (t Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	fmt.Println("QueryRowContext")
	callerName := utils.GetFnCallerName(1, 2)
	ctx = context.WithValue(
		ctx,
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)
	t.doBefore(ctx, query, args...)

	res := t.tx.QueryRowContext(ctx, query, args...)
	err := res.Err()

	defer t.doAfter(ctx, err, query, args...)

	return res
}

func (t Tx) ExecContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (res sql.Result, err error) {
	callerName := utils.GetFnCallerName(1, 2)
	ctx = context.WithValue(
		ctx,
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)
	t.doBefore(ctx, query, args...)
	defer t.doAfter(ctx, err, query, args...)

	res, err = t.tx.ExecContext(ctx, query, args...)

	return
}

func (t Tx) Rollback() error {
	return t.tx.Rollback()
}

func (t Tx) Commit() error {
	return t.tx.Commit()
}

func (t Tx) doBefore(ctx context.Context, query string, args ...interface{}) {
	for _, f := range t.before {
		f(ctx, query, args...)
	}
}

func (t Tx) doAfter(ctx context.Context, err error, query string, args ...interface{}) {
	for _, f := range t.after {
		f(ctx, err, query, args...)
	}
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		DB:     db,
		before: make([]BeforeFunc, 0),
		after:  make([]AfterFunc, 0),
	}
}

func (db *DB) AddBeforeFunc(f BeforeFunc) {
	db.before = append(db.before, f)
}

func (db *DB) AddAfterFunc(f AfterFunc) {
	db.after = append(db.after, f)
}

func (db DB) doBefore(ctx context.Context, query string, args ...interface{}) context.Context {
	for _, f := range db.before {
		ctx = f(ctx, query, args...)
	}

	return ctx
}

func (db DB) doAfter(ctx context.Context, err error, query string, args ...interface{}) {
	for _, f := range db.after {
		f(ctx, err, query, args...)
	}
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

func (db DB) Exec(query string, args ...interface{}) (res sql.Result, err error) {
	callerName := utils.GetFnCallerName(1, 2)
	ctx := context.WithValue(
		context.Background(),
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)

	ctx = db.doBefore(ctx, query, args...)
	defer db.doAfter(ctx, err, query, args...)

	res, err = db.DB.Exec(query, args...)

	return
}

func (db DB) ExecContext(c context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	ctx, span := telemetry.NewDatabaseSpan(c, query)
	defer span.End()

	callerName := utils.GetFnCallerName(1, 2)
	ctx = context.WithValue(
		ctx,
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)

	ctx = db.doBefore(ctx, query, args...)
	defer db.doAfter(ctx, err, query, args...)

	res, err = TxFromContext(ctx, db.DB).ExecContext(ctx, query, args...)

	return
}

func (db DB) Query(query string, args ...interface{}) (res *sql.Rows, err error) {
	callerName := utils.GetFnCallerName(1, 2)
	ctx := context.WithValue(
		context.Background(),
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)

	ctx = db.doBefore(ctx, query, args...)
	defer db.doAfter(ctx, err, query, args...)

	res, err = db.DB.Query(query, args...)

	return
}

func (db DB) QueryContext(c context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	ctx, span := telemetry.NewDatabaseSpan(c, query)
	defer span.End()

	callerName := utils.GetFnCallerName(1, 2)
	ctx = context.WithValue(
		ctx,
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)

	ctx = db.doBefore(ctx, query, args...)
	defer db.doAfter(ctx, err, query, args...)

	rows, err = TxFromContext(ctx, db.DB).QueryContext(ctx, query, args...)

	return
}

func (db DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}

func (db DB) QueryRowContext(c context.Context, query string, args ...interface{}) *sql.Row {
	ctx, span := telemetry.NewDatabaseSpan(c, query)
	defer span.End()

	callerName := utils.GetFnCallerName(1, 2)
	ctx = context.WithValue(
		ctx,
		constants.ContextKeySqlWrapperCaller,
		callerName,
	)

	ctx = db.doBefore(ctx, query, args...)

	res := TxFromContext(ctx, db.DB).QueryRowContext(ctx, query, args...)
	err := res.Err()

	defer db.doAfter(ctx, err, query, args...)

	return res
}
