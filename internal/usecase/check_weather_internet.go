package usecase

import (
	"errors"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"strconv"
)

type TemperatureDTO struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type CheckWeatherUseCase struct {
	WeatherapiIntegration entity.WeatherapiIntegrationInterface
	ViacepIntegration     entity.ViacepIntegrationInterface
}

func NewCheckWeatherUseCase(
	weatherapiIntegration entity.WeatherapiIntegrationInterface,
	viacepIntegration entity.ViacepIntegrationInterface) *CheckWeatherUseCase {

	return &CheckWeatherUseCase{
		WeatherapiIntegration: weatherapiIntegration,
		ViacepIntegration:     viacepIntegration,
	}
}

func cepIsValid(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	_, err := strconv.Atoi(cep)
	if err != nil {
		return false
	}
	return true
}

func (c *CheckWeatherUseCase) Execute(cep string) (*TemperatureDTO, error) {
	if !cepIsValid(cep) {
		return &TemperatureDTO{}, errors.New("invalid zipcode")
	}
	chViacep := make(chan types.Either[string])
	go c.ViacepIntegration.GetCity(cep, chViacep)
	reschViacep := <-chViacep
	if reschViacep.Left != nil {
		return &TemperatureDTO{}, reschViacep.Left
	}
	chWeatherApi := make(chan types.Either[float64])
	go c.WeatherapiIntegration.GetCelsiusTemperatureByCity(reschViacep.Right, chWeatherApi)
	reschWeatherApi := <-chWeatherApi
	if reschWeatherApi.Left != nil {
		return &TemperatureDTO{}, reschWeatherApi.Left
	}

	ent := entity.NewWeather(reschWeatherApi.Right)

	dto := TemperatureDTO{
		TempCelsius:    ent.TemperatureCelsius,
		TempFahrenheit: ent.TemperatureFahrenheit,
		TempKelvin:     ent.TemperatureKelvin,
	}
	return &dto, nil
}
