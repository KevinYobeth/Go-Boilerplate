package repository

import (
	"context"

	eventcontract "github.com/kevinyobeth/go-boilerplate/internal/shared/event_contract"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/event"
	"github.com/ztrue/tracerr"
)

type RabbitMQAuthenticationPublisher struct {
	publisher event.PublisherInterface
}

func NewRabbitMQAuthenticationPublisher(publisher event.PublisherInterface) Publisher {
	return &RabbitMQAuthenticationPublisher{publisher: publisher}
}

func (p *RabbitMQAuthenticationPublisher) UserRegistered(c context.Context, payload eventcontract.UserRegistered) error {
	err := p.publisher.Publish(c, event.Event{
		Event: eventcontract.UserRegisteredEvent,
		Data:  payload,
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
