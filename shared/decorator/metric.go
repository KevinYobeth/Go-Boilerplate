package decorator

import (
	"context"
	"fmt"
	"go-boilerplate/shared/metrics"
	"strings"
	"time"

	"github.com/samber/lo"
)

type commandMetricDecorator[C any] struct {
	base   CommandHandler[C]
	client metrics.Client
}

func (d commandMetricDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	start := time.Now()

	actionName := generateMetricName(cmd)

	defer func() {
		end := time.Since(start)

		d.client.DurationMs(fmt.Sprintf("commands_%s_duration_ms", actionName), int(end.Milliseconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("commands_%s_success", actionName), 1)
			return
		}

		d.client.Inc(fmt.Sprintf("commands_%s_failure", actionName), 1)
	}()

	return d.base.Handle(ctx, cmd)
}

type queryMetricDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	client metrics.Client
}

func (d queryMetricDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	start := time.Now()

	actionName := generateMetricName(cmd)

	defer func() {
		end := time.Since(start)

		d.client.DurationMs(fmt.Sprintf("queries_%s_duration_ms", actionName), int(end.Milliseconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("queries_%s_success", actionName), 1)
			return
		}

		d.client.Inc(fmt.Sprintf("queries_%s_failure", actionName), 1)
	}()

	return d.base.Handle(ctx, cmd)
}

func generateMetricName(handler any) string {
	parts := strings.Split(fmt.Sprintf("%T", handler), ".")

	return lo.SnakeCase(parts[1])
}
