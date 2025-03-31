package utils

import (
	"context"
	"go-boilerplate/shared/constants"
)

func GetRequestIDFromContext(ctx context.Context) string {
	return ctx.Value(constants.ContextKeyRequestID).(string)
}

func GetTraceIDFromContext(ctx context.Context) string {
	return ctx.Value(constants.ContextKeyTraceID).(string)
}