package job

import (
	"context"

	"github.com/go-co-op/gocron/v2"
	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/command"
)

func (j Job) RegisterUpdateLinkVisitSnapshotJob(ctx context.Context) {
	jobName := "UpdateLinkVisitSnapshotJob"
	config := config.LoadSettingConfig()

	j.schedule.NewJob(gocron.DurationJob(config.SettingLinkVisitSnapshotInterval), gocron.NewTask(
		func() {
			j.decorate(jobName, func() error {
				return j.app.Commands.UpdateLinkVisitSnapshot.Handle(ctx, &command.UpdateLinkVisitSnapshotRequest{})
			})
		},
	), gocron.WithSingletonMode(gocron.LimitModeReschedule))
}
