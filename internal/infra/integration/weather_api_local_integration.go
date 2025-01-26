package integration

import (
	"bytes"
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"io"
	"log"
)

type WeatherAPILocalIntegration struct {
	uri string
}

func NewWeatherAPILocalIntegration(uri string) *WeatherAPILocalIntegration {
	return &WeatherAPILocalIntegration{
		uri: uri,
	}
}

func (w *WeatherAPILocalIntegration) GetCep(cep *dto.WeatherPostDTO, resultCh chan<- types.Either[dto.TemperatureDTO]) {
	client := getHttpClient()
	weatherUrl := w.uri

	filter, err := json.Marshal(cep)
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}

	req, err := client.Post(weatherUrl, "application/json", bytes.NewReader(filter))
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body")
		}
	}(req.Body)
	res, err := io.ReadAll(req.Body)
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	var data dto.TemperatureDTO
	err = json.Unmarshal(res, &data)
	if err != nil {
		resultCh <- types.Either[dto.TemperatureDTO]{Left: err}
		return
	}
	resultCh <- types.Either[dto.TemperatureDTO]{Right: data}
	close(resultCh)
}
