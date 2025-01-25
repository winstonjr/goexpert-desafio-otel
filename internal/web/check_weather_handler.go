package web

import (
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
	"log"
	"net/http"
)

type WeatherPostInternalHandler struct {
	checkWeatherLocalUseCase entity.CheckWeatherLocalUseCaseInterface
}

func NewWeatherPostInternalHandler(checkWeatherLocalUseCase entity.CheckWeatherLocalUseCaseInterface) *WeatherPostInternalHandler {
	return &WeatherPostInternalHandler{
		checkWeatherLocalUseCase: checkWeatherLocalUseCase,
	}
}

func (wph *WeatherPostInternalHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var wp dto.WeatherPostDTO
	err := json.NewDecoder(r.Body).Decode(&wp)
	if err != nil || wp.CEP == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`invalid zipcode`))
		return
	}
	log.Println("cep acquired", wp.CEP)
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	w.WriteHeader(http.StatusUnprocessableEntity)
	//	_, _ = w.Write([]byte(`invalid zipcode`))
	//	return
	//}
	//log.Println("cep acquired", string(body))
	temperature, err := wph.checkWeatherLocalUseCase.ExecuteLocal(&wp)
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
