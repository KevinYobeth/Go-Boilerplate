package telemetry

import (
	"context"
	"net"
	"runtime"
	"strings"
	"time"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"

	"github.com/redis/go-redis/extra/rediscmd/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type RedisTracingHook struct {
}

func NewRedisTracingHook() redis.Hook {
	return RedisTracingHook{}
}

func (hook RedisTracingHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (hook RedisTracingHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		if cmd.Name() == "ping" {
			return next(ctx, cmd)
		}

		ctx, span := NewRedisSpan(ctx, cmd)
		defer span.End()

		return next(ctx, cmd)
	}
}

func (hook RedisTracingHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		ctx, span := NewRedisSpan(ctx, cmds[0])
		defer span.End()

		return next(ctx, cmds)
	}
}

func NewRedisSpan(c context.Context, cmd redis.Cmder) (context.Context, trace.Span) {
	now := time.Now().UTC()

	actionName := generateRedisActionName()

	ctx, span := StartSpan(c, actionName,
		trace.WithSpanKind(trace.SpanKindInternal),
	)

	span.SetAttributes(
		attribute.String("db.timestamp", now.String()),
		attribute.String("db.statement", rediscmd.CmdString(cmd)),
		attribute.String(RequestIDKey, utils.GetRequestIDFromContext(ctx)),
	)

	return ctx, span
}

func generateRedisActionName() string {
	file, _, _ := funcFileLine("github.com/redis/go-redis")

	if file != "" {
		return file
	}

	return "repository.cache.unknown"
}

func funcFileLine(pkg string) (string, string, int) {
	const depth = 16
	var pcs [depth]uintptr
	n := runtime.Callers(5, pcs[:])
	ff := runtime.CallersFrames(pcs[:n])

	var fn, file string
	var line int
	for {
		f, ok := ff.Next()
		if !ok {
			break
		}
		fn, file, line = f.Function, f.File, f.Line
		if !strings.Contains(fn, pkg) {
			break
		}
	}

	if ind := strings.LastIndexByte(fn, '/'); ind != -1 {
		fn = fn[ind+1:]
	}

	return fn, file, line
}
