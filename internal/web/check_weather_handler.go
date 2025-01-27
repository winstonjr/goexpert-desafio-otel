package web

import (
	"context"
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
)

type WeatherPostInternalHandler struct {
	context                  *context.Context
	checkWeatherLocalUseCase entity.CheckWeatherLocalUseCaseInterface
	OTELTracer               trace.Tracer
}

func NewWeatherPostInternalHandler(checkWeatherLocalUseCase entity.CheckWeatherLocalUseCaseInterface, tracer trace.Tracer) *WeatherPostInternalHandler {
	return &WeatherPostInternalHandler{
		checkWeatherLocalUseCase: checkWeatherLocalUseCase,
		OTELTracer:               tracer,
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

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, spanInicial := wph.OTELTracer.Start(ctx, "SPAN INICIAL: entrance-cep-api")
	spanInicial.End()
	ctx, span := wph.OTELTracer.Start(ctx, "CHAMADA EXTERNA: entrance-cep-api")
	defer span.End()

	temperature, err := wph.checkWeatherLocalUseCase.ExecuteLocal(ctx, &wp)
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
