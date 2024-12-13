package entity

type Weather struct {
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}

func (w *Weather) CalculateFahrenheit() float64 {
	w.TemperatureFahrenheit = (w.TemperatureCelsius * 1.8) + 32
	return w.TemperatureFahrenheit
}

func (w *Weather) CalculateKelvin() float64 {
	w.TemperatureKelvin = w.TemperatureCelsius + 273
	return w.TemperatureKelvin
}

func NewWeather(celsius float64) *Weather {
	w := &Weather{TemperatureCelsius: celsius}
	w.CalculateFahrenheit()
	w.CalculateKelvin()
	return w
}
