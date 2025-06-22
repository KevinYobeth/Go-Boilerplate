package telemetry

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewIntraprocessSpan(ctx context.Context) (context.Context, trace.Span) {
	now := time.Now().UTC()

	spanName := generateIntraprocessActionName()

	ctx, span := StartSpan(ctx, spanName,
		trace.WithSpanKind(trace.SpanKindInternal),
	)

	span.SetAttributes(
		attribute.Stringer("db.timestamp", now),
		attribute.String(RequestIDKey, utils.GetRequestIDFromContext(ctx)),
	)

	return ctx, span
}

func generateIntraprocessActionName() string {
	pc, _, _, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fnName := strings.Split(fn.Name(), ".")

		return fmt.Sprintf("%s.%s.%s", "intraprocess", fnName[1], fnName[2])
	}
	return "intraprocess.unknown"
}
