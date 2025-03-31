package decorator

import (
	"context"
	"encoding/json"
	"fmt"
	"go-boilerplate/shared/utils"
	"strings"

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

	d.logger = d.logger.With(fields...)
	d.logger.Debug("executing command")

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

	d.logger = d.logger.With(fields...)
	d.logger.Debug("executing query")

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
	var body any = fmt.Sprintf("%#v", param)

	jsonBytes, err := json.Marshal(param)
	if err == nil {
		var param map[string]any
		if err := json.Unmarshal(jsonBytes, &param); err == nil {
			body = param
		}
	}

	return body
}
