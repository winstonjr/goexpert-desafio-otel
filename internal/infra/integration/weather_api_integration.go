package integration

import (
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"io"
	"net/url"
)

type WeatherapiIntegration struct {
	apiKey string
}

func NewWeatherApiIntegration(apiKey string) *WeatherapiIntegration {
	return &WeatherapiIntegration{apiKey: apiKey}
}

func (w *WeatherapiIntegration) GetCelsiusTemperatureByCity(city string, resultch chan<- types.Either[float64]) {
	client := getHttpClient()
	weatherUrl := "https://api.weatherapi.com/v1/current.json?key=" + w.apiKey + "&q=" + url.QueryEscape(city) + "&aqi=no"
	req, err := client.Get(weatherUrl)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
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
