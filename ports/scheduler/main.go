package scheduler

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinyobeth/go-boilerplate/internal/link/presentation/job"
	linkService "github.com/kevinyobeth/go-boilerplate/internal/link/services"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/constants"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/graceroutine"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/log"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"

	"github.com/go-co-op/gocron/v2"
)

func RunScheduler() {
	logger := log.InitLogger()
	s, err := gocron.NewScheduler()

	if err != nil {
		logger.Errorf("Failed to create scheduler: %v", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	linkServices := linkService.NewLinkService()
	linkJob := job.NewJob(s, linkServices, logger)

	ctx := context.Background()
	c := context.WithValue(ctx, constants.ContextKeyRequestID, utils.RandomString(10))

	linkJob.Run(c)

	<-signals

	logger.Info("Shutting down scheduler...")
	if err := s.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Scheduler Shutdown")
}
