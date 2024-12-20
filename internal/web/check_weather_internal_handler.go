package web

import (
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
	"log"
	"net/http"
)

type WeatherPostHandler struct {
	checkWeatherUseCase entity.CheckWeatherUseCaseInterface
}

type WeatherPostDTO struct {
	CEP string
}

func NewWeatherPostHandler(checkWeatherUseCase entity.CheckWeatherUseCaseInterface) *WeatherPostHandler {
	return &WeatherPostHandler{
		checkWeatherUseCase: checkWeatherUseCase,
	}
}

func (wph *WeatherPostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var dto WeatherPostDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil || dto.CEP == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`invalid zipcode`))
		return
	}
	log.Println("cep acquired", dto.CEP)
	temperature, err := wph.checkWeatherUseCase.Execute(dto.CEP)
	if err != nil {
		log.Println("temperature error", err.Error())
		if err.Error() == "invalid zipcode" {
			log.Println("temperature error", 422)
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = w.Write([]byte(`invalid zipcode`))
			return
		} else if err.Error() == "can not find zipcode" {
			log.Println("temperature error", 404)
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`can not find zipcode`))
			return
		} else {
			log.Println("temperature error", 500)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = json.NewEncoder(w).Encode(temperature)
	if err != nil {
		log.Println("write to buffer error", 500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
