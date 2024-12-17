package usecase

import (
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
)

type CheckWeatherLocalUseCase struct {
	LocalApiIntegration entity.CheckWeatherLocalUseCaseInterface
}

func (c *CheckWeatherLocalUseCase) ExecuteLocal(cep string) (*dto.TemperatureDTO, error) {
	//	if !cepIsValid(cep) {
	//		return nil, errors.New("invalid zipcode")
	//	}
	//	chViaCep := make(chan types.Either[dto.TemperatureDTO])
	//	go c.LocalApiIntegration.GetCity(cep, chViaCep)
	//	resChViaCep := <-chViaCep
	//	if resChViaCep.Left != nil {
	//		return nil, resChViaCep.Left
	//	}
	//	chWeatherApi := make(chan types.Either[float64])
	//	go c.WeatherApiIntegration.GetCelsiusTemperatureByCity(resChViaCep.Right, chWeatherApi)
	//	resChWeatherApi := <-chWeatherApi
	//	if resChWeatherApi.Left != nil {
	//		return nil, resChWeatherApi.Left
	//	}
	//
	//	ent := entity.NewWeather(resChWeatherApi.Right)

	return &dto.TemperatureDTO{
		City:           "",
		TempCelsius:    0,
		TempFahrenheit: 0,
		TempKelvin:     0,
	}, nil

}
