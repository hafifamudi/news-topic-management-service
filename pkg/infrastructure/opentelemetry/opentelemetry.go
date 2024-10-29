package opentelemetry

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"news-topic-management-service/pkg/infrastructure/tracer/jeager"
)

func InitOpenTelemetry() (func(context.Context) error, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create Prometheus exporter: %w", err)
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(exporter))
	otel.SetMeterProvider(meterProvider)

	tp, err := jeager.StartTracing()
	if err != nil {
		return nil, fmt.Errorf("failed to start tracing: %w", err)
	}
	otel.SetTracerProvider(tp)

	return func(ctx context.Context) error {
		if err := tp.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	}, nil
}
