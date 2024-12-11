package integration

import (
	"encoding/json"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"io"
	"net/url"
)

type locationDTO struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type conditionDTO struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type currentDTO struct {
	LastUpdatedEpoch int          `json:"last_updated_epoch"`
	LastUpdated      string       `json:"last_updated"`
	TempC            float64      `json:"temp_c"`
	TempF            float64      `json:"temp_f"`
	IsDay            int          `json:"is_day"`
	Condition        conditionDTO `json:"condition"`
	WindMph          float64      `json:"wind_mph"`
	WindKph          float64      `json:"wind_kph"`
	WindDegree       int          `json:"wind_degree"`
	WindDir          string       `json:"wind_dir"`
	PressureMb       float64      `json:"pressure_mb"`
	PressureIn       float64      `json:"pressure_in"`
	PrecipMm         float64      `json:"precip_mm"`
	PrecipIn         float64      `json:"precip_in"`
	Humidity         int          `json:"humidity"`
	Cloud            int          `json:"cloud"`
	FeelslikeC       float64      `json:"feelslike_c"`
	FeelslikeF       float64      `json:"feelslike_f"`
	WindchillC       float64      `json:"windchill_c"`
	WindchillF       float64      `json:"windchill_f"`
	HeatindexC       float64      `json:"heatindex_c"`
	HeatindexF       float64      `json:"heatindex_f"`
	DewpointC        float64      `json:"dewpoint_c"`
	DewpointF        float64      `json:"dewpoint_f"`
	VisKm            float64      `json:"vis_km"`
	VisMiles         float64      `json:"vis_miles"`
	Uv               float64      `json:"uv"`
	GustMph          float64      `json:"gust_mph"`
	GustKph          float64      `json:"gust_kph"`
}

type weatherDTO struct {
	Location locationDTO `json:"location"`
	Current  currentDTO  `json:"current"`
}

type WeatherapiIntegration struct {
	apiKey string
}

func NewWeatherapiIntegration(apiKey string) *WeatherapiIntegration {
	return &WeatherapiIntegration{apiKey: apiKey}
}

func (w *WeatherapiIntegration) GetCelsiusTemperatureByCity(city string, resultch chan<- types.Either[float64]) {
	client := getHttpClient()
	weatherUrl := "https://api.weatherapi.com/v1/current.json?key=" + w.apiKey + "&q=" + url.QueryEscape(city) + "&aqi=no"
	req, err := client.Get(weatherUrl)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	var data weatherDTO
	err = json.Unmarshal(res, &data)
	if err != nil {
		resultch <- types.Either[float64]{Left: err}
		return
	}
	resultch <- types.Either[float64]{Right: data.Current.TempC}
	close(resultch)
}
