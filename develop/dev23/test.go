package main

import "fmt"

/*

=== Задача ===

Удалить i-ый элемент из слайса.

*/

func main() {
	// Удалить элемент можно двумя способами

	// 1-ый быстрый, но его применять имеет смысл, если порядок элементов в слайсе не важен
	args := []int{5, 2, 13, 19, 4, 9}
	idxForDel := 2

	args[idxForDel] = args[len(args)-1]
	args = args[:len(args)-1]

	fmt.Printf("Измененный слайс: %v\n", args)

	// 2-ый медленне, но его применять имеет смысл, если порядок элементов в слайсе важен
	args = []int{5, 2, 13, 19, 4, 9}
	idxForDel = 3

	args = append(args[:idxForDel], args[idxForDel+1:]...)
	fmt.Printf("Измененный слайс: %v\n", args)
}
