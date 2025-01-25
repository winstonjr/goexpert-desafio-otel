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
	GetCep(cep *dto.WeatherPostDTO, resultCh chan<- types.Either[dto.TemperatureDTO])
}

type CheckWeatherUseCaseInterface interface {
	Execute(cep string) (*dto.TemperatureDTO, error)
}

type CheckWeatherLocalUseCaseInterface interface {
	ExecuteLocal(cep *dto.WeatherPostDTO) (*dto.TemperatureDTO, error)
}
