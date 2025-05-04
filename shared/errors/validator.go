package errors

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewValidatorError(err validator.ValidationErrorsTranslations) error {
	metadata := make(map[string]string)

	for key, value := range err {
		metadata[strings.Split(key, ".")[1]] = value
	}

	return NewIncorrectInputWithMetadataError(nil, "validation error", metadata)
}
