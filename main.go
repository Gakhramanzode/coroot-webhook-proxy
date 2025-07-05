package main

import (
	"fmt"
	"log"
	"net/http"

	"context"

	"coroot-webhook-proxy/config"
	"coroot-webhook-proxy/handler"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() {
	ctx := context.Background()
	client := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(ctx, client)

	if err != nil {
		log.Fatalf("failed to initialize exporter: %e", err)
	}

	res, err := resource.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize resource: %e", err)
	}

	// Create the trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// Set the global trace provider
	otel.SetTracerProvider(tp)

	// Set the propagator
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(propagator)
}

func main() {
	// Initialize OpenTelemetry tracer
	initTracer()

	// Initialize the HTTP server with instrumentation
	router := http.NewServeMux()

	cfg := config.LoadConfig()

	router.Handle("/health", otelhttp.NewHandler(http.HandlerFunc(handler.HealthHandler), "GET /health"))
	router.Handle("/coroot-webhook", otelhttp.NewHandler(http.HandlerFunc(handler.WebhookHandler(cfg)), "POST /coroot-webhook"))

	fmt.Println("Server is running on :8080")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
