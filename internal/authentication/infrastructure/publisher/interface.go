package publisher

import (
	"context"

	eventcontract "github.com/kevinyobeth/go-boilerplate/internal/shared/event_contract"
)

type Publisher interface {
	UserRegistered(c context.Context, payload eventcontract.UserRegistered) error
}
