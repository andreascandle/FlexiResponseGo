package observability

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	trc "go.opentelemetry.io/otel/trace"
)

var globalTracer trc.Tracer

// InitTracer initializes the OpenTelemetry tracer with custom exporter URL support.
func InitTracer(serviceName, exporterURL string) func() {
	exporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint(exporterURL))
	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tracerProvider)

	globalTracer = otel.Tracer(serviceName)

	return func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}
}

// StartSpan starts a new span in the context with attributes.
func StartSpan(ctx context.Context, spanName string, attributes map[string]string) (context.Context, trc.Span) {
	opts := []trc.SpanStartOption{}
	for k, v := range attributes {
		opts = append(opts, trc.WithAttributes(attribute.String(k, v)))
	}
	return globalTracer.Start(ctx, spanName, opts...)
}
