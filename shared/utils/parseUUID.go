package utils

import (
	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedUUID, nil
}
