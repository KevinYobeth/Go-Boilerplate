package scheduler

import (
	"go-boilerplate/shared/graceroutine"
	"go-boilerplate/shared/log"
	"go-boilerplate/src/books/presentation/job"
	"go-boilerplate/src/books/services"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-co-op/gocron/v2"
)

func RunScheduler() {
	logger := log.InitLogger()
	app := services.NewBookService()

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
