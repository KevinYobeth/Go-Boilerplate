package metrics

import (
	"context"
	"go-boilerplate/shared/telemetry"

	"go.opentelemetry.io/otel/metric"
)

type metricsClient struct {
	meter metric.Meter
}

type Client interface {
	Inc(key string, value int)
	DurationMs(key string, value int)
}

func InitClient() Client {
	meter := telemetry.GetMetric()

	return &metricsClient{
		meter: meter,
	}
}

func (c *metricsClient) Inc(key string, value int) {
	counter, _ := c.meter.Int64Counter(key)
	counter.Add(context.Background(), int64(value))
}

func (c *metricsClient) DurationMs(key string, value int) {
	histogram, _ := c.meter.Int64Histogram(key, metric.WithUnit("ms"))
	histogram.Record(context.Background(), int64(value))
}
