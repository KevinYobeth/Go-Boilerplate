package event

import (
	"bytes"
	"encoding/json"
)

type Event struct {
	Event string
	Data  any
}

func NewEvent(event string, data any) Event {
	return Event{event, data}
}

func Serialize(event Event) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(event)
	return b.Bytes(), err
}

func Deserialize(b []byte) (Event, error) {
	var event Event
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&event)
	return event, err
}

func (e Event) TransformTo(data interface{}) error {
	b, err := json.Marshal(e.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		return err
	}

	return nil
}
