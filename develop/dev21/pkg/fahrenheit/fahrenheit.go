package fahrenheit

// Представим, что это, некая внешняя библиотека, которая работает только с градусами Фаренгейта
type fahrenheit struct {
	temperature float64
}

type FahrenheitTemperature interface {
	GetTemperature() float64
}

func (f *fahrenheit) GetTemperature() float64 {
	return f.temperature
}

// Создаем конструктор
func NewFahrenheit(temperature float64) FahrenheitTemperature {
	return &fahrenheit{
		temperature: temperature,
	}
}
