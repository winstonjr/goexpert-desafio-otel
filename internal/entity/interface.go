package entity

import (
	"context"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
)

type ViacepIntegrationInterface interface {
	GetCity(ctx context.Context, cep string, resultch chan<- types.Either[string])
}

type WeatherapiIntegrationInterface interface {
	GetCelsiusTemperatureByCity(ctx context.Context, city string, resultch chan<- types.Either[float64])
}

type WeatherApiLocalIntegrationInterface interface {
	GetCep(ctx context.Context, cep *dto.WeatherPostDTO, resultCh chan<- types.Either[dto.TemperatureDTO])
}

type CheckWeatherUseCaseInterface interface {
	Execute(ctx context.Context, cep string) (*dto.TemperatureDTO, error)
}

type CheckWeatherLocalUseCaseInterface interface {
	ExecuteLocal(ctx context.Context, cep *dto.WeatherPostDTO) (*dto.TemperatureDTO, error)
}
