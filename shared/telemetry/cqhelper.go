package telemetry

import (
	"context"
	"fmt"
	"go-boilerplate/shared/utils"
	"runtime"
	"strings"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewCQHelperSpan(ctx context.Context) (context.Context, trace.Span) {
	now := time.Now().UTC()

	actionName := generateCQHelperActionName()

	ctx, span := StartSpan(ctx, actionName,
		trace.WithSpanKind(trace.SpanKindInternal),
	)

	span.SetAttributes(
		attribute.Stringer("cqrs.helper.timestamp", now),
		attribute.String("cqrs.type", "helper"),
		attribute.String("cqrs.operation", actionName),
		attribute.Stringer("cqrs.timestamp", now),
		attribute.String(RequestIDKey, utils.GetRequestIDFromContext(ctx)),
	)

	return ctx, span
}

func generateCQHelperActionName() string {
	pc, _, _, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fnName := strings.Split(fn.Name(), ".")

		return fmt.Sprintf("%s.%s", "helper", fnName[1])
	}
	return "helper.unknown"
}
