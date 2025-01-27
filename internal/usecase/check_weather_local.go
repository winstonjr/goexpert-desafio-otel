package usecase

import (
	"context"
	"errors"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
)

type CheckWeatherLocalUseCase struct {
	LocalApiIntegration entity.WeatherApiLocalIntegrationInterface
}

func NewCheckWeatherLocalUseCase(localIntegration entity.WeatherApiLocalIntegrationInterface) *CheckWeatherLocalUseCase {
	return &CheckWeatherLocalUseCase{
		LocalApiIntegration: localIntegration,
	}
}

func (c *CheckWeatherLocalUseCase) ExecuteLocal(ctx context.Context, cep *dto.WeatherPostDTO) (*dto.TemperatureDTO, error) {
	if !cepIsValid(cep.CEP) {
		return nil, errors.New("invalid zipcode")
	}
	chLocal := make(chan types.Either[dto.TemperatureDTO])
	go c.LocalApiIntegration.GetCep(ctx, cep, chLocal)
	resChLocal := <-chLocal
	if resChLocal.Left != nil {
		return nil, resChLocal.Left
	}

	return &resChLocal.Right, nil

}
