package scheduler

import (
	"os"
	"os/signal"
	"syscall"

	authorIntraprocess "github.com/kevinyobeth/go-boilerplate/internal/authors/presentation/intraprocess"
	authorsService "github.com/kevinyobeth/go-boilerplate/internal/authors/services"
	"github.com/kevinyobeth/go-boilerplate/internal/books/infrastructure/intraprocess"
	"github.com/kevinyobeth/go-boilerplate/internal/books/presentation/job"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services"
	"github.com/kevinyobeth/go-boilerplate/shared/graceroutine"
	"github.com/kevinyobeth/go-boilerplate/shared/log"

	"github.com/go-co-op/gocron/v2"
)

func RunScheduler() {
	logger := log.InitLogger()
	authorsService := authorsService.NewAuthorService()
	authorIntraprocess := authorIntraprocess.NewAuthorIntraprocessService(authorsService)

	booksAuthorIntraprocess := intraprocess.NewBookAuthorIntraprocessService(authorIntraprocess)

	app := services.NewBookService(booksAuthorIntraprocess)

	s, err := gocron.NewScheduler()

	if err != nil {
		logger.Errorf("Failed to create scheduler: %v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	jobs := job.NewJob(s, app, logger)

	jobs.Run()

	<-signals

	logger.Info("Shutting down scheduler...")
	if err := s.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Scheduler Shutdown")
}
