package usecase

import (
	"errors"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"strconv"
)

type CheckWeatherUseCase struct {
	WeatherApiIntegration entity.WeatherapiIntegrationInterface
	ViaCepIntegration     entity.ViacepIntegrationInterface
}

func NewCheckWeatherUseCase(
	weatherApiIntegration entity.WeatherapiIntegrationInterface,
	viaCepIntegration entity.ViacepIntegrationInterface) *CheckWeatherUseCase {

	return &CheckWeatherUseCase{
		WeatherApiIntegration: weatherApiIntegration,
		ViaCepIntegration:     viaCepIntegration,
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

func (c *CheckWeatherUseCase) Execute(cep string) (*entity.TemperatureDTO, error) {
	if !cepIsValid(cep) {
		return nil, errors.New("invalid zipcode")
	}
	chViaCep := make(chan types.Either[string])
	go c.ViaCepIntegration.GetCity(cep, chViaCep)
	resChViaCep := <-chViaCep
	if resChViaCep.Left != nil {
		return nil, resChViaCep.Left
	}
	chWeatherApi := make(chan types.Either[float64])
	go c.WeatherApiIntegration.GetCelsiusTemperatureByCity(resChViaCep.Right, chWeatherApi)
	reschWeatherApi := <-chWeatherApi
	if reschWeatherApi.Left != nil {
		return nil, reschWeatherApi.Left
	}

	ent := entity.NewWeather(reschWeatherApi.Right)

	dto := entity.TemperatureDTO{
		City:           resChViaCep.Right,
		TempCelsius:    ent.TemperatureCelsius,
		TempFahrenheit: ent.TemperatureFahrenheit,
		TempKelvin:     ent.TemperatureKelvin,
	}
	return &dto, nil
}
