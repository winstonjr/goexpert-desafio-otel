package dto

type TemperatureDTO struct {
	City           string  `json:"city"`
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type ViacepDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type LocationDTO struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type ConditionDTO struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type CurrentDTO struct {
	LastUpdatedEpoch int          `json:"last_updated_epoch"`
	LastUpdated      string       `json:"last_updated"`
	TempC            float64      `json:"temp_c"`
	TempF            float64      `json:"temp_f"`
	IsDay            int          `json:"is_day"`
	Condition        ConditionDTO `json:"condition"`
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

type WeatherDTO struct {
	Location LocationDTO `json:"location"`
	Current  CurrentDTO  `json:"current"`
}

type WeatherPostDTO struct {
	CEP string
}
