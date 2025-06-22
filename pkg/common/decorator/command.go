package decorator

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"go.uber.org/zap"
)

func ApplyCommandDecorators[H any](
	handler CommandHandler[H],
	logger *zap.SugaredLogger,
	metricsClient metrics.Client,
) CommandHandler[H] {
	return commandOTelDecorator[H]{
		commandErrorDecorator[H]{
			commandLoggingDecorator[H]{
				commandMetricDecorator[H]{
					base:   handler,
					client: metricsClient,
				},
				logger,
			},
		},
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
