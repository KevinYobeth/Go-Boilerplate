package scheduler

import (
	"go-boilerplate/shared/log"
	"go-boilerplate/src/books/job"
	"go-boilerplate/src/books/services"

	"github.com/go-co-op/gocron/v2"
)

func RunScheduler() {
	logger := log.InitLogger()
	app := services.NewBookService()

	s, err := gocron.NewScheduler()

	if err != nil {
		logger.Errorf("Failed to create scheduler: %v", err)
	}

	jobs := job.NewJob(s, app, logger)

	jobs.Run()

	var forever chan struct{}
	<-forever

	if err != nil {
		logger.Errorf("Failed to create job: %v", err)
	}
}
