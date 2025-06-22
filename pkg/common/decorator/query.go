package decorator

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"

	"go.uber.org/zap"
)

func ApplyQueryDecorators[H any, R any](
	handler QueryHandler[H, R],
	logger *zap.SugaredLogger,
	metricsClient metrics.Client,
) QueryHandler[H, R] {
	return queryOTelDecorator[H, R]{
		queryErrorDecorator[H, R]{
			queryLoggingDecorator[H, R]{
				queryMetricDecorator[H, R]{
					base:   handler,
					client: metricsClient,
				},
				logger,
			},
		},
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
