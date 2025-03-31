package decorator

import (
	"context"

	"go.uber.org/zap"
)

func ApplyQueryDecorators[H any, R any](
	handler QueryHandler[H, R],
	logger *zap.SugaredLogger,
) QueryHandler[H, R] {
	return queryOTelDecorator[H, R]{
		queryLoggingDecorator[H, R]{
			base:   handler,
			logger: logger,
		},
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
