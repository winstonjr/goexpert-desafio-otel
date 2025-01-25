package main

import (
	"fmt"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/integration"
	"github.com/winstonjr/goexpert-desafio-otel/internal/usecase"
	"github.com/winstonjr/goexpert-desafio-otel/internal/web"
	"net/http"
)

func main() {
	localApiIntegration := integration.NewWeatherAPILocalIntegration()
	localApiUseCase := usecase.NewCheckWeatherLocalUseCase(localApiIntegration)

	localApiPostHandler := web.NewWeatherPostInternalHandler(localApiUseCase)

	http.HandleFunc("/", localApiPostHandler.Handle)

	fmt.Println("Service B - Listening on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
