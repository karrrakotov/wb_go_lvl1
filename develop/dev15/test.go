package main

import "fmt"

/*

=== Задача ===

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

var justString string

func main() {
	// Вызываем функцию
	someFunc()
}

func someFunc() {
	str := createHugeString(1 << 10)

	// Для безопасности, выполняем проверку на размер строки
	if len(str) >= 100 {
		justString = str[:100]
		fmt.Printf("Длина строки была больше 100, новая строка равна: %v\n", len(justString))
	} else {
		fmt.Printf("Длина строки меньше 100, она равна: %v\n", len(str))
	}
}

// Создаем функцию для создания большой строки
func createHugeString(size int) string {
	b := make([]byte, size)
	s := string(b)
	return s
}
