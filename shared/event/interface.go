package event

import "context"

type PublisherInterface interface {
	Publish(event Event) error
}

type SubscriberInterface interface {
	Subscribe(c context.Context, handler func(context.Context, Event) error) error
}
