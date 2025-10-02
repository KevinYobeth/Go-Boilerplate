package utils

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/shared/constants"
)

func GetRequestIDFromContext(ctx context.Context) string {
	val := ReadFromCtx(ctx, constants.ContextKeyRequestID)
	if val == nil {
		return ""
	}

	return val.(string)
}

func GetTraceIDFromContext(ctx context.Context) string {
	val := ReadFromCtx(ctx, constants.ContextKeyTraceID)
	if val == nil {
		return ""
	}

	return val.(string)
}

func GetSpanIDFromContext(ctx context.Context) string {
	val := ReadFromCtx(ctx, constants.ContextKeySpanID)
	if val == nil {
		return ""
	}

	return val.(string)
}
