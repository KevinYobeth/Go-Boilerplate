package decorator

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

func ApplyCommandDecorators[H any](
	handler CommandHandler[H],
	logger *zap.SugaredLogger,
) CommandHandler[H] {
	return handler
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
