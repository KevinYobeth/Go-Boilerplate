package pagination

import (
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
)

func ValidateLimitPaginationParams(page, limit *uint64) error {
	if page != nil && *page < 1 {
		return errors.NewIncorrectInputError(nil, "page must be greater than or equal to 1")
	}

	if limit != nil && *limit < 1 {
		return errors.NewIncorrectInputError(nil, "limit must be greater than or equal to 1")
	}
	return nil
}
