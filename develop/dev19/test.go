package main

import "fmt"

/*

=== Задача ===

Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.

*/

func main() {
	str := "главрыба"
	reverseStr := reverse(str)
	fmt.Printf("%v - %v\n", str, reverseStr)

	// Проверка, где байты расположены в памяти
	str = "Hello, 今天快 樂的一天"
	for i, c := range str {
		fmt.Printf("Символ %c начинается на позиции: %d\n", c, i)
	}

	// Проверка работы reverse со строкой выше
	reverseStr = reverse(str)
	fmt.Printf("%v - %v\n", str, reverseStr)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
