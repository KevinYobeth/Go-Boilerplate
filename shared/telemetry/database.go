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

func NewDatabaseSpan(c context.Context, query string) (context.Context, trace.Span) {
	now := time.Now().UTC()

	actionName := generateDatabaseActionName()

	ctx, span := StartSpan(c, actionName,
		trace.WithSpanKind(trace.SpanKindInternal),
	)

	span.SetAttributes(
		attribute.String("db.statement", query),
		attribute.Stringer("db.timestamp", now),
		attribute.String(RequestIDKey, utils.GetRequestIDFromContext(ctx)),
	)

	return ctx, span
}

func generateDatabaseActionName() string {
	pc, _, _, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fnName := strings.Split(fn.Name(), ".")

		return fmt.Sprintf("%s.%s.%s", "repository", fnName[1], fnName[2])
	}
	return "repository.unknown"
}
