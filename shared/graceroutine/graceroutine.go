package graceroutine

import (
	"context"
	"sync"

	"github.com/panjf2000/ants"
)

var (
	wg                    sync.WaitGroup
	antsPool, _           = ants.NewPool(ants.DEFAULT_ANTS_POOL_SIZE)
	globalCtx, cancelFunc = context.WithCancel(context.Background())
)

// SetPool set pool
func SetPool(p *ants.Pool) {
	antsPool = p
}

// Context get graceroutine ctx
func Context() context.Context {
	return globalCtx
}

// Submit submit and run task on goroutine
func Submit(task func()) error {
	wg.Add(1)

	err := antsPool.Submit(func() {
		defer wg.Done()
		task()
	})

	if err != nil {
		wg.Done()
	}

	return err
}

// MustSubmit submit and run task on goroutine
func MustSubmit(task func()) {
	err := Submit(task)
	if err != nil {
		panic(err)
	}
}

// Stop call cancel func
func Stop() {
	cancelFunc()
}

// Wait wait default WaitGroup. Recommended to run this on main function
func Wait() {
	wg.Wait()
}
