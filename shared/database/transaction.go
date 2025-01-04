package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type TransactionManager struct {
	db PostgresDB
}

func NewTransactionManager(db PostgresDB) TransactionManager {
	return TransactionManager{db}
}

func (tm TransactionManager) RunInTransaction(c context.Context, f func(c context.Context) error) error {
	var err error

	tx := TransactionFromContext(c)

	if tx == nil {
		tx, err = tm.db.BeginTx(c, nil)
		if err != nil {
			return fmt.Errorf("failed to start transaction: %w", err)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			if tx != nil {
				errRollback := tx.Rollback()
				if errRollback != nil {
					err = errRollback
				}
			}

			panic(r)
		}
	}()

	ctx := ContextWithTx(c, tx)

	err = f(ctx)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

type ctxKeyType int

const txKey ctxKeyType = 0

// TxFromContext returns the transaction object from the context. Return db if not exists
func TxFromContext(ctx context.Context, db sq.StdSqlCtx) Queryer {
	tx, ok := ctx.Value(txKey).(Transaction)
	if !ok {
		return db
	}

	return tx
}

// TransactionFromContext returns the transaction object from the context. Return db if not exists
func TransactionFromContext(ctx context.Context) Transaction {
	tx, ok := ctx.Value(txKey).(Transaction)
	if !ok {
		return nil
	}

	return tx
}

// ContextWithTx add database transaction to context
func ContextWithTx(parentContext context.Context, tx Transaction) context.Context {
	return context.WithValue(parentContext, txKey, tx)
}
