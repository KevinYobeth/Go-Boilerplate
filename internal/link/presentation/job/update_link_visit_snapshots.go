package job

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/command"
)

func (j Job) RegisterUpdateLinkVisitSnapshotJob(ctx context.Context) {
	jobName := "UpdateLinkVisitSnapshotJob"

	j.schedule.NewJob(gocron.DurationJob(5*time.Second), gocron.NewTask(
		func() {
			j.decorate(jobName, func() error {
				return j.app.Commands.UpdateLinkVisitSnapshot.Handle(ctx, &command.UpdateLinkVisitSnapshotRequest{})
			})
		},
	), gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
}
