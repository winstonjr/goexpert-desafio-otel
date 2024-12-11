package entity

import "github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"

type ViacepIntegrationInterface interface {
	GetCity(cep string, resultch chan<- types.Either[string])
}

type WeatherapiIntegrationInterface interface {
	GetCelsiusTemperatureByCity(city string, resultch chan<- types.Either[float64])
}
