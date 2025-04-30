package telemetry

import (
	"context"
	"errors"
	"go-boilerplate/config"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	logTrace "go.opentelemetry.io/otel/sdk/log"
	metricTrace "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	RequestIDKey = "request.id"
	TraceIDKey   = "trace.id"
	SpanIDKey    = "span.id"
)

func InitOtel(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	tracerProvider, err := newTracerProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	otel.SetTracerProvider(tracerProvider)

	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)

	otel.SetMeterProvider(meterProvider)

	loggerProvider, err := newLoggerProvider()
	if err != nil {
		handleErr(err)
		return
	}

	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)

	global.SetLoggerProvider(loggerProvider)

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func Resource() *resource.Resource {
	cfg := config.LoadAppConfig()

	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfg.AppName),
		semconv.ServiceVersionKey.String(cfg.AppVersion),
	)
}

func newMeterProvider() (*metricTrace.MeterProvider, error) {
	cfg := config.LoadOpenTelemetryConfig()

	exporter, err := otlpmetricgrpc.New(context.Background(),
		otlpmetricgrpc.WithEndpointURL(cfg.OtelGRPCEndpoint),
		otlpmetricgrpc.WithRetry(otlpmetricgrpc.RetryConfig{
			Enabled:         cfg.OtelRetryEnabled,
			InitialInterval: cfg.OtelRetryInitialInterval,
			MaxInterval:     cfg.OtelRetryMaxInterval,
			MaxElapsedTime:  cfg.OtelRetryMaxElapsedTime,
		}),
	)
	if err != nil {
		return nil, err
	}

	meterProvider := metricTrace.NewMeterProvider(
		metricTrace.WithReader(
			metricTrace.NewPeriodicReader(exporter,
				metricTrace.WithInterval(5*time.Second),
			),
		),
		metricTrace.WithResource(Resource()),
	)

	return meterProvider, nil
}

func newLoggerProvider() (*logTrace.LoggerProvider, error) {
	cfg := config.LoadOpenTelemetryConfig()

	exporter, err := otlploggrpc.New(context.Background(),
		otlploggrpc.WithEndpointURL(cfg.OtelGRPCEndpoint),
		otlploggrpc.WithRetry(otlploggrpc.RetryConfig{
			Enabled:         cfg.OtelRetryEnabled,
			InitialInterval: cfg.OtelRetryInitialInterval,
			MaxInterval:     cfg.OtelRetryMaxInterval,
			MaxElapsedTime:  cfg.OtelRetryMaxElapsedTime,
		}),
	)
	if err != nil {
		return nil, err
	}

	logProvider := logTrace.NewLoggerProvider(
		logTrace.WithProcessor(
			logTrace.NewBatchProcessor(exporter,
				logTrace.WithExportTimeout(5*time.Second),
			),
		),
		logTrace.WithResource(Resource()),
	)

	return logProvider, nil
}

func newTracerProvider() (*sdkTrace.TracerProvider, error) {
	cfg := config.LoadOpenTelemetryConfig()

	if cfg.OtelDisabled {
		return sdkTrace.NewTracerProvider(
			sdkTrace.WithResource(Resource()),
		), nil
	}

	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithEndpointURL(cfg.OtelGRPCEndpoint),
		otlptracegrpc.WithRetry(otlptracegrpc.RetryConfig{
			Enabled:         cfg.OtelRetryEnabled,
			InitialInterval: cfg.OtelRetryInitialInterval,
			MaxInterval:     cfg.OtelRetryMaxInterval,
			MaxElapsedTime:  cfg.OtelRetryMaxElapsedTime,
		}),
	)
	if err != nil {
		return nil, err
	}

	tracerProvider := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(exporter,
			sdkTrace.WithBatchTimeout(5*time.Second),
		),
		sdkTrace.WithResource(Resource()),
	)

	return tracerProvider, nil
}

func GetTracer() trace.Tracer {
	cfg := config.LoadAppConfig()

	tracer := otel.GetTracerProvider().Tracer(cfg.AppName,
		trace.WithInstrumentationVersion(cfg.AppVersion))
	return tracer
}

func GetMetric() metric.Meter {
	cfg := config.LoadAppConfig()

	meter := otel.GetMeterProvider().Meter(cfg.AppName,
		metric.WithInstrumentationVersion(cfg.AppVersion),
	)
	return meter
}

func StartSpan(
	ctx context.Context,
	spanName string,
	opts ...trace.SpanStartOption,
) (context.Context, trace.Span) {
	tracer := GetTracer()

	opts = append([]trace.SpanStartOption{
		trace.WithTimestamp(time.Now().UTC()),
	}, opts...)

	return tracer.Start(ctx, spanName, opts...)
}

func GetTraceID(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return ""
	}

	return span.SpanContext().TraceID().String()
}

func GetSpanID(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return ""
	}

	return span.SpanContext().SpanID().String()
}
