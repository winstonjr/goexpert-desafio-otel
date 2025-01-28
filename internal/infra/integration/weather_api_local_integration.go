package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"log"
	"net/http"
)

type WeatherAPILocalIntegration struct {
	uri string
}

func NewWeatherAPILocalIntegration(uri string) *WeatherAPILocalIntegration {
	return &WeatherAPILocalIntegration{
		uri: uri,
	}
}

func (w *WeatherAPILocalIntegration) GetCep(ctx context.Context, cep *dto.WeatherPostDTO, resultCh chan<- types.Either[dto.TemperatureDTO]) {
	client := getHttpClient()
	weatherUrl := w.uri

	filter, err := json.Marshal(cep)
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", weatherUrl, bytes.NewReader(filter))
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := client.Do(req)
	//resp, err := client.Post(weatherUrl, "application/json", bytes.NewReader(filter))
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body")
		}
	}(resp.Body)
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	var data dto.TemperatureDTO
	err = json.Unmarshal(res, &data)
	if err != nil {
		errAttempt := string(res)
		if errAttempt != "" {
			resultCh <- types.Either[dto.TemperatureDTO]{Left: errors.New(errAttempt)}
			return
		}
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	resultCh <- types.Either[dto.TemperatureDTO]{Right: data}
	close(resultCh)
}
