package entity

import (
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
)

type ViacepIntegrationInterface interface {
	GetCity(cep string, resultch chan<- types.Either[string])
}

type WeatherapiIntegrationInterface interface {
	GetCelsiusTemperatureByCity(city string, resultch chan<- types.Either[float64])
}

type WeatherApiLocalIntegrationInterface interface {
	GetCep(cep string, resultCh chan<- types.Either[float64])
}

type CheckWeatherUseCaseInterface interface {
	Execute(cep string) (*dto.TemperatureDTO, error)
}

type CheckWeatherLocalUseCaseInterface interface {
	ExecuteLocal(cep string) (*dto.TemperatureDTO, error)
}
