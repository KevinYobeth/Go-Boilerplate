package constants

type ContextKey string

const (
	ContextKeyRequestID ContextKey = "request_id"
	ContextKeyTraceID   ContextKey = "trace_id"
	ContextKeySpanID    ContextKey = "span_id"

	ContextKeyClaims ContextKey = "claims"
)
