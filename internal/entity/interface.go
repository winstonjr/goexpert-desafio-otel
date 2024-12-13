package entity

import (
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
)

type TemperatureDTO struct {
	City           string  `json:"city"`
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type ViacepIntegrationInterface interface {
	GetCity(cep string, resultch chan<- types.Either[string])
}

type WeatherapiIntegrationInterface interface {
	GetCelsiusTemperatureByCity(city string, resultch chan<- types.Either[float64])
}

type CheckWeatherUseCaseInterface interface {
	Execute(cep string) (*TemperatureDTO, error)
}
