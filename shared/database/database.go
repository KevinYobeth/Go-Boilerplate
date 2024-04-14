package database

import (
	"context"
	"fmt"
)

type TransactionManager struct {
	db PostgresDB
}

func NewTransactionManager(db PostgresDB) *TransactionManager {
	return &TransactionManager{db}
}

func (tm *TransactionManager) RunInTransaction(c context.Context, f func(c context.Context) error) error {
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
