package pagination

import (
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
)

func ValidateLimitPaginationParams(page, limit int) error {
	if page != 0 && page < 1 {
		return errors.NewIncorrectInputError(nil, "page must be greater than or equal to 1")
	}

	if limit != 0 && limit < 1 {
		return errors.NewIncorrectInputError(nil, "limit must be greater than or equal to 1")
	}
	return nil
}
