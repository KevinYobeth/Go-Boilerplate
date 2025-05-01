package utils

import (
	"github.com/kevinyobeth/go-boilerplate/shared/errors"

	"github.com/google/uuid"
)

func ParseUUID(id string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, errors.NewIncorrectInputError(err, "invalid UUID")
	}

	return parsedUUID, nil
}
