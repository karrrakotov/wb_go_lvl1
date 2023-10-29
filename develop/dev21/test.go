package main

/*

=== Задача ===

Реализовать паттерн «адаптер» на любом примере.

*/

import (
	"fmt"

	"karrrakotov/wb_go_lvl1/dev21/pkg/fahrenheit"
)

// Представим, что наш проект работает только с градусами по Цельсию
// И у нас есть, некий внешний пакет (pkg/fahrenheit), который работает только с Фаренгейтами
// А нас это не устраивает и нам нужен Адаптер

// Наша основная структура
type celsius struct {
	temperature float64
}

type CelsiusTemperature interface {
	GetTemperature() float64
}

func (c *celsius) GetTemperature() float64 {
	return c.temperature
}

func (c *celsius) SetTemperature(temperature float64) {
	c.temperature = temperature
}

// Создаем Адаптер, который будет дружить Цельсии с Фаренгейтами
type fahrenheitAdapter struct {
	fahrenheit fahrenheit.FahrenheitTemperature
}

func (f *fahrenheitAdapter) GetTemperature() float64 {
	return (f.fahrenheit.GetTemperature() - 32) * 5 / 9
}

func main() {
	fahrenheit := fahrenheit.NewFahrenheit(50.0)
	fahrenheitAdapter := NewFahrenheitAdapter(fahrenheit)
	fmt.Printf("Полученное значение Цельсия через адатер: %v\n", fahrenheitAdapter.GetTemperature())

	// Устанавливаем полученное значение с Fahrenheit в наш Celsius
	celsius := NewCelsius(fahrenheitAdapter.GetTemperature())
	fmt.Printf("Наше значение в цельсиях: %v\n", celsius.GetTemperature())
}

// Создаем конструкторы
func NewCelsius(temperature float64) CelsiusTemperature {
	return &celsius{
		temperature: temperature,
	}
}

func NewFahrenheitAdapter(fahrenheit fahrenheit.FahrenheitTemperature) *fahrenheitAdapter {
	return &fahrenheitAdapter{
		fahrenheit: fahrenheit,
	}
}
