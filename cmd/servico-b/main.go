package main

import (
	"context"
	"fmt"
	"github.com/winstonjr/goexpert-desafio-otel/configs"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/integration"
	"github.com/winstonjr/goexpert-desafio-otel/internal/usecase"
	"github.com/winstonjr/goexpert-desafio-otel/internal/web"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	// TODO: colocar estas informações em variáveis de ambiente
	shutdown, err := initProvider("entrance-cep-api", "otel-collector:4317")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()
	tracer := otel.Tracer("microservice-tracer")

	viaCepIntegration := integration.NewViacepIntegration()
	weatherApiIntegration := integration.NewWeatherApiIntegration(config.WeatherApiKey)
	checkWeatherUseCase := usecase.NewCheckWeatherUseCase(weatherApiIntegration, viaCepIntegration)

	weatherPostHandler := web.NewWeatherPostHandler(checkWeatherUseCase, tracer)

	http.HandleFunc("/", weatherPostHandler.Handle)

	fmt.Println("Service B - Listening on port :8081")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func initProvider(serviceName, collectorURL string) (func(ctx context.Context) error, error) {
	ctx := context.Background()
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	//conn, err := grpc.NewClient(collectorURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, collectorURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to collector: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}
