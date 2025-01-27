package integration

import (
	"context"
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
	"net/url"
)

type WeatherapiIntegration struct {
	apiKey string
}

func NewWeatherApiIntegration(apiKey string) *WeatherapiIntegration {
	return &WeatherapiIntegration{apiKey: apiKey}
}

func (w *WeatherapiIntegration) GetCelsiusTemperatureByCity(ctx context.Context, city string, resultch chan<- types.Either[float64]) {
	client := getHttpClient()
	weatherUrl := "https://api.weatherapi.com/v1/current.json?key=" + w.apiKey + "&q=" + url.QueryEscape(city) + "&aqi=no"

	//resp, err := client.Get(weatherUrl)
	req, err := http.NewRequestWithContext(ctx, "GET", weatherUrl, nil)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := client.Do(req)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	var data dto.WeatherDTO
	err = json.Unmarshal(res, &data)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	resultch <- types.Either[float64]{Right: data.Current.TempC}
	close(resultch)
}
