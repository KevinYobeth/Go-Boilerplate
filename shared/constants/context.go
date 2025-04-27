package constants

type ContextKey string

const (
	ContextKeyRequestID ContextKey = "request_id"
	ContextKeyTraceID   ContextKey = "trace_id"

	ContextKeyClaims ContextKey = "claims"
)
