package event

import "context"

type Event struct {
	c    context.Context
	data any
}

func NewEvent(c context.Context, data any) Event {
	return Event{c, data}
}
