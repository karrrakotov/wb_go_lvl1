package main

import (
	"fmt"
	"math/big"
)

/*

=== Задача ===

Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

*/

func main() {
	// Для большинства операция подойдет и вместимость int64, если нужно еще больше, можно использовать uint64(для положительных), либо же использовать float64
	// Но если числовые значения слишком огромные, то необходимо использовать пакет math/big

	a := "24e20"
	b := "24e15"

	// Инициализация объектов типа big
	num1 := new(big.Float)
	num2 := new(big.Float)

	// Преобразование строковых значений в big
	_, success1 := num1.SetString(a)
	_, success2 := num2.SetString(b)

	if success1 && success2 {
		// Умножение
		resultMul := new(big.Float).Mul(num1, num2)

		// Деление
		resultDev := new(big.Float).Quo(num1, num2)

		// Сложение
		resultSum := new(big.Float).Add(num1, num2)

		// Вычитание
		resultSub := new(big.Float).Sub(num1, num2)
		fmt.Printf("Результаты:\nУмножение: %v\nДеление: %v\nСложение: %v\nВычитание: %v\n", resultMul, resultDev, resultSum, resultSub)
	} else {
		fmt.Println("Ошибка преобразования строки в число")
	}
}
