package decorator

import (
	"context"

	"go.uber.org/zap"
)

func ApplyCommandDecorators[H any](
	handler CommandHandler[H],
	logger *zap.SugaredLogger,
) CommandHandler[H] {
	return commandOTelDecorator[H]{
		commandLoggingDecorator[H]{
			base:   handler,
			logger: logger,
		},
	}
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
