package event

import "context"

type PublisherInterface interface {
	Publish(c context.Context, event Event) error
}

type SubscriberInterface interface {
	Subscribe(c context.Context, handler func(context.Context, Event) error) error
	Shutdown() error
}
