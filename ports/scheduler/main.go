package scheduler

import (
	authorIntraprocess "go-boilerplate/internal/authors/presentation/intraprocess"
	authorsService "go-boilerplate/internal/authors/services"
	"go-boilerplate/internal/books/infrastructure/intraprocess"
	"go-boilerplate/internal/books/presentation/job"
	"go-boilerplate/internal/books/services"
	"go-boilerplate/shared/graceroutine"
	"go-boilerplate/shared/log"
	"os"
	"os/signal"
	"syscall"

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
