package main

/*

=== Задача ===

Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

*/

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

type Service interface {
	Distance(p1, p2 *point) float64
}

// Функция для вычисления расстояния между двумя точками
func (p *point) Distance(another *point) float64 {
	return math.Sqrt(math.Pow(another.x-p.x, 2) + math.Pow(another.y-p.y, 2))
}

func main() {
	point1, point2 := NewPoint(1, 3), NewPoint(4, 6)
	distance := point1.Distance(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}

// Конструктор
func NewPoint(x, y float64) *point {
	return &point{
		x: x,
		y: y,
	}
}
