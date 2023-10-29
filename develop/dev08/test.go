package main

import "fmt"

/*

=== Задача ===

Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

*/

func main() {
	var number uint = 6
	var bit uint = 0
	var pos uint = 1

	// Вариант 1
	fmt.Printf("Изначальное число: %v\n", number)
	newNum := changeBit(number, bit, pos)
	fmt.Printf("Новое число после установки бита: %v\n", newNum)

	// Вариант 2
	fmt.Printf("Изначальное число: %v\n", number)
	newNum2 := toggleBit(number, pos)
	fmt.Printf("Новое число после установки бита: %v\n", newNum2)
}

// Вариант 1
func changeBit(number, bit, pos uint) (newNum uint) {
	var mask uint = 1 << pos

	if bit == 1 {
		newNum = number | mask // установка бита в 1
	} else if bit == 0 {
		newNum = number &^ mask // установка бита в 0
	} else {
		fmt.Printf("Ошибка, неверно указан bit: %v", bit)
		return
	}

	return newNum
}

// Вариант 2
func toggleBit(n, pos uint) uint {
	return n ^ (1 << pos)
}
