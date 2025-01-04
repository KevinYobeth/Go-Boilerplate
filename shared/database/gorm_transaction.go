package database

import (
	"context"
	"fmt"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type GormTransactionManager struct {
	db *gorm.DB
}

func GlobalTransactionCallback(db *gorm.DB) {
	ctx := db.Statement.Context
	if ctx == nil {
		return
	}

	fmt.Println("running callback")
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		fmt.Println("tx found")
		// db.InstanceSet("tx", tx)
		db.Statement.ConnPool = tx.ConnPool
		return
	} else {
		fmt.Println("tx not found")
		db.Begin()
		fmt.Println("tx begin from callback")
	}
}

func NewGormTransactionManager(db *gorm.DB) GormTransactionManager {
	return GormTransactionManager{db}
}

func (tm GormTransactionManager) RunInTransaction(c context.Context, f func(c context.Context) error) error {
	var err error

	tx := GormTransactionFromContext(c)

	if tx == nil {
		tx = tm.db.Begin()
		fmt.Println("tx begin")
	}

	defer func() {
		if r := recover(); r != nil {
			if tx != nil {
				err := tx.Rollback().Error
				fmt.Println("tx rollback")
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	ctx := GormContextWithTx(c, tx)
	err = f(ctx)
	if err != nil {
		errRollback := tx.Rollback().Error
		fmt.Println("tx rollback 2", errRollback)
		if errRollback != nil {
			return tracerr.Wrap(errRollback)
		}

		return tracerr.Wrap(err)
	}

	err = tx.Commit().Error
	fmt.Println("tx commit")
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// TxFromContext returns the transaction object from the context. Return db if not exists
func GormTxFromContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		return db
	}

	return tx
}

func GormTransactionFromContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		return nil
	}

	return tx
}

func GormContextWithTx(parentContext context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(parentContext, txKey, tx)
}
