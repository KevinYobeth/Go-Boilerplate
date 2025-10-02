package event

import (
	"encoding/json"
)

type Event struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

func (e Event) TransformTo(data any) error {
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
