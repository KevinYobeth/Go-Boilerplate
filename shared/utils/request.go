package utils

import (
	"context"
	"go-boilerplate/shared/constants"
)

func GetRequestIDFromContext(ctx context.Context) string {
	return ReadFromCtx(ctx, constants.ContextKeyRequestID).(string)
}

func GetTraceIDFromContext(ctx context.Context) string {
	return ReadFromCtx(ctx, constants.ContextKeyTraceID).(string)
}

func GetSpanIDFromContext(ctx context.Context) string {
	return ReadFromCtx(ctx, constants.ContextKeySpanID).(string)
}
