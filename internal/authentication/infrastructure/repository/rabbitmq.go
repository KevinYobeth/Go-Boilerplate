package repository

import (
	"context"

	interfaces "github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces/event"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/event"
	"github.com/ztrue/tracerr"
)

type RabbitMQAuthenticationPublisher struct {
	publisher event.PublisherInterface
}

func NewRabbitMQAuthenticationPublisher(publisher event.PublisherInterface) Publisher {
	return &RabbitMQAuthenticationPublisher{publisher: publisher}
}

func (p *RabbitMQAuthenticationPublisher) UserRegistered(c context.Context, payload interfaces.UserRegistered) error {
	err := p.publisher.Publish(c, event.Event{
		Event: interfaces.UserRegisteredEvent,
		Data:  payload,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
