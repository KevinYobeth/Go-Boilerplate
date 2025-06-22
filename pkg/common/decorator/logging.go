package decorator

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/log"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/utils"
	"go.uber.org/zap"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *zap.SugaredLogger
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	fields := []interface{}{
		"command", generateActionName(cmd),
		"command_body", getLogBodyParam(cmd),
	}

	if reqID := utils.GetRequestIDFromContext(ctx); reqID != "" {
		fields = append(fields, "request_id", reqID)
	}

	logger := log.WithTrace(ctx, d.logger).With(fields...)

	logger.Debug("executing command")

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[Q any, R any] struct {
	base   QueryHandler[Q, R]
	logger *zap.SugaredLogger
}

func (d queryLoggingDecorator[Q, R]) Handle(ctx context.Context, qry Q) (result R, err error) {
	fields := []interface{}{
		"query", generateActionName(qry),
		"query_body", getLogBodyParam(qry),
	}

	if reqID := utils.GetRequestIDFromContext(ctx); reqID != "" {
		fields = append(fields, "request_id", reqID)
	}

	logger := log.WithTrace(ctx, d.logger).With(fields...)

	logger.Debug("executing command")

	return d.base.Handle(ctx, qry)
}

func generateActionName(handler any) string {
	parts := strings.Split(fmt.Sprintf("%T", handler), ".")
	if len(parts) < 2 {
		return "cqrs.unknown"
	}
	action := parts[1]
	return fmt.Sprintf("%s.%s", "cqrs", action)
}

func getLogBodyParam(param any) any {
	redactedFields := []string{
		"password",
	}

	var body any = fmt.Sprintf("%#v", param)

	jsonBytes, err := json.Marshal(param)
	if err == nil {
		var param map[string]any
		if err := json.Unmarshal(jsonBytes, &param); err == nil {
			body = param
		}
	}

	for _, field := range redactedFields {
		if _, ok := body.(map[string]any)[field]; ok {
			body.(map[string]any)[field] = "************"
		}
	}

	return body
}
