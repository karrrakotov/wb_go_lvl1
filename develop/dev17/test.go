package main

import "fmt"

/*

=== Задача ===

Реализовать бинарный поиск встроенными методами языка.

*/

func main() {
	args := []int{-5, 2, 4, 7, 9, 13, 13, 16, 22, 25, 33}

	position := binarySearch(args, -9)
	fmt.Printf("Индекс искомого числа: %v\n", position)
}

func binarySearch(args []int, findNum int) (pos int) {
	left, right := 0, len(args)-1

	for left <= right {
		middle := left + (right-left)/2

		if args[middle] == findNum {
			pos = middle
			return pos
		} else if args[middle] < findNum {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
