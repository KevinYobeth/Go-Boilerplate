package job

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services"
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

func (j Job) Run(ctx context.Context) {
	j.RegisterUpdateLinkVisitSnapshotJob(ctx)

	j.schedule.Start()
}

func (j Job) decorate(name string, fn func() error) {
	j.logger.Infof("The job %s has been started", name)

	if err := fn(); err != nil {
		j.logger.Errorf("Error encountered while running the job: %v", err)
	}

	j.logger.Infof("The job %s has been finished", name)
}
