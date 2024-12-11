package entity

type Weather struct {
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}

func (w *Weather) CalculateFahrenheit() {
	w.TemperatureFahrenheit = (w.TemperatureCelsius * 1.8) + 32
}

func (w *Weather) CalculateKelvin() {
	w.TemperatureKelvin = w.TemperatureCelsius + 273
}

func NewWeather(celsius float64) *Weather {
	w := &Weather{TemperatureCelsius: celsius}
	w.CalculateFahrenheit()
	w.CalculateKelvin()
	return w
}
