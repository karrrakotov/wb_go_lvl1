package main

import "fmt"

/*

=== Задача ===

Поменять местами два числа без создания временной переменной.

*/

func main() {
	val1, val2 := 2, 5

	fmt.Printf("val1 = %v, val2 = %v\n", val1, val2)
	val1 = val1 + val2 // 7
	val2 = val1 - val2 // 7 - 5 = 2
	val1 = val1 - val2 // 7 - 2 = 5
	fmt.Println("Поменяли местами..")
	fmt.Printf("val1 = %v, val2 = %v\n", val1, val2)
}
