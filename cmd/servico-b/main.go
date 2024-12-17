package main

import (
	"fmt"
	"github.com/winstonjr/goexpert-desafio-otel/configs"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/integration"
	"github.com/winstonjr/goexpert-desafio-otel/internal/usecase"
	"github.com/winstonjr/goexpert-desafio-otel/internal/web"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	viaCepIntegration := integration.NewViacepIntegration()
	weatherApiIntegration := integration.NewWeatherApiIntegration(config.WeatherApiKey)
	checkWeatherUseCase := usecase.NewCheckWeatherUseCase(weatherApiIntegration, viaCepIntegration)

	weatherPostHandler := web.NewWeatherPostHandler(checkWeatherUseCase)

	http.HandleFunc("/", weatherPostHandler.Handle)

	fmt.Println("Service B - Listening on port :8081")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
