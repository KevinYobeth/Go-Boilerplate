package telemetry

import (
	"context"
	"go-boilerplate/shared/utils"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(ctx context.Context, name string, kv ...attribute.KeyValue) (context.Context, trace.Span) {
	now := time.Now().UTC()

	ctx, span := StartSpan(ctx, name,
		trace.WithSpanKind(trace.SpanKindInternal),
	)

	attributes := []attribute.KeyValue{
		attribute.Stringer("cqrs.timestamp", now),
		attribute.String(RequestIDKey, utils.GetRequestIDFromContext(ctx)),
	}

	attributes = append(attributes, kv...)

	span.SetAttributes(
		attributes...,
	)

	return ctx, span
}
