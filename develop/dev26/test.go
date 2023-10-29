package main

/*

=== Задача ===

Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false

*/

import (
	"fmt"
	"sync"
	"unicode"
)

func main() {
	args := []string{"abcd", "abCdefAaf", "aabcd"}
	var wg sync.WaitGroup

	// Подрубаем горутины для быстренького выполнения задачки
	for _, val := range args {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			res := checkUniq(val)
			fmt.Printf("%v - %v\n", val, res)
		}(val)
	}

	wg.Wait()
}

// Создаем функцию для проверки на уникальность
func checkUniq(str string) (res bool) {
	checkMap := make(map[rune]bool)

	for _, char := range str {
		lowChar := unicode.ToUpper(char)

		if checkMap[lowChar] {
			return
		}

		checkMap[lowChar] = true
	}
	return true
}
