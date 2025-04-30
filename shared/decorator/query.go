package decorator

import (
	"context"
	"go-boilerplate/shared/metrics"

	"go.uber.org/zap"
)

func ApplyQueryDecorators[H any, R any](
	handler QueryHandler[H, R],
	logger *zap.SugaredLogger,
	metricsClient metrics.Client,
) QueryHandler[H, R] {
	return queryOTelDecorator[H, R]{
		queryLoggingDecorator[H, R]{
			queryMetricDecorator[H, R]{
				base:   handler,
				client: metricsClient,
			},
			logger,
		},
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
