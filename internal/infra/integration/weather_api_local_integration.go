package integration

import (
	"bytes"
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"io"
	"log"
)

type WeatherAPILocalIntegration struct{}

func (w *WeatherAPILocalIntegration) GetCep(cep string, resultCh chan<- types.Either[dto.TemperatureDTO]) {
	client := getHttpClient()
	weatherUrl := "http://localhost:8081/"
	req, err := client.Post(weatherUrl, "application/json", bytes.NewBufferString(cep))
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
