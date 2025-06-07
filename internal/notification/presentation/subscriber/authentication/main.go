package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinyobeth/go-boilerplate/internal/notification/services"
	"github.com/kevinyobeth/go-boilerplate/internal/notification/services/command"
	interfaces "github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces/event"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/topic"
	"github.com/kevinyobeth/go-boilerplate/shared/event"
	"github.com/kevinyobeth/go-boilerplate/shared/graceroutine"
	"github.com/kevinyobeth/go-boilerplate/shared/log"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
)

type HandlerParams struct {
	logger *zap.SugaredLogger
	app    *services.Application
}

func main() {
	logger := log.InitLogger()
	c := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	subscriber := event.InitSubscriber(event.SubscriberOptions{Topic: topic.AuthenticationTopic})
	app := services.NewNotificationService()

	go func() {
		subscriber.Subscribe(c, func(ctx context.Context, e event.Event) error {
			logger.Infof("Received event: %s", e.Event)
			var err error

			switch e.Event {
			case interfaces.UserRegisteredEvent:
				err = onUserRegistered(ctx, HandlerParams{logger, &app}, e)
			default:
				logger.Infof("Event %s is not handled", e.Event)
			}

			if err != nil {
				logger.Errorf("Error handling event %s: %v", e.Event, tracerr.Wrap(err))
			} else {
				logger.Infof("Successfully handled event: %s", e.Event)
			}

			return tracerr.Wrap(err)
		})
	}()

	<-signals

	logger.Info("Shutting down subscriber...")

	if err := subscriber.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	graceroutine.Stop()
	graceroutine.Wait()

	logger.Info("Subscriber Shutdown")
}

func onUserRegistered(c context.Context, params HandlerParams, e event.Event) error {
	var data interfaces.UserRegistered
	err := e.TransformTo(&data)

	if err != nil {
		params.logger.Errorf("Failed to transform event data: %v", tracerr.Wrap(err))
	}

	err = params.app.Commands.SendWelcomeNotification.Handle(c, &command.SendWelcomeNotificationRequest{
		UserID: data.UserID,
		Name:   data.Name,
		Email:  data.Email,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
