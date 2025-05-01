package job

import (
	"context"
	"time"

	"github.com/kevinyobeth/go-boilerplate/internal/books/services"
	"github.com/kevinyobeth/go-boilerplate/internal/books/services/command"

	"github.com/go-co-op/gocron/v2"
	"go.uber.org/zap"
)

type Job struct {
	schedule gocron.Scheduler
	app      services.Application
	logger   *zap.SugaredLogger
}

func NewJob(schedule gocron.Scheduler, app services.Application, logger *zap.SugaredLogger) Job {
	return Job{schedule, app, logger}
}

func (j Job) RegisterAutomaticArchiveBooksJob() {
	jobName := "Automatic Archive Books"

	j.schedule.NewJob(gocron.DurationJob(5*time.Second), gocron.NewTask(
		func() {
			j.decorate(jobName, func() error {
				j.app.Commands.CreateBook.Handle(context.Background(), command.CreateBookRequest{
					Title:  "Automatic Archive Book",
					Author: "Automatic Archive Author",
				})
				return nil
			})
		},
	),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
}

func (j Job) decorate(name string, fn func() error) {
	j.logger.Infof("The job %s has been started", name)

	if err := fn(); err != nil {
		j.logger.Errorf("Error encountered while running the job: %v", err)
	}

	j.logger.Infof("The job %s has been finished", name)
}

func (j Job) Run() {
	j.RegisterAutomaticArchiveBooksJob()

	j.schedule.Start()
}
