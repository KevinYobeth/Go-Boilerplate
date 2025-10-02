package rabbitmq

import (
	"bytes"
	"encoding/json"

	"github.com/kevinyobeth/go-boilerplate/shared/event"
)

func NewEvent(eventKey string, data any) event.Event {
	return event.Event{Event: eventKey, Data: data}
}

func Serialize(event event.Event) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(event)
	return b.Bytes(), err
}

func Deserialize(b []byte) (event.Event, error) {
	var event event.Event
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&event)
	return event, err
}
