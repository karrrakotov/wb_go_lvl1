package main

import (
	"fmt"
)

/*

=== Задача ===

Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

*/

// В данной задаче, были рассмотрены все виды сортировок, для более глубокого понимания
func main() {
	// Сортировка пузырьком
	args := []int{2, 1, 9, 0, 4, -2}
	fmt.Printf("До сортировки: %v\n", args)
	bubbleSort(args)
	fmt.Printf("После сортировки: %v\n", args)

	// Сортировка вставками
	args = []int{2, 1, 0}
	fmt.Printf("\nДо сортировки: %v\n", args)
	insertionSort(args)
	fmt.Printf("После сортировки: %v\n", args)

	// Сортировка слиянием
	args = []int{21, 1, 22, 10, -5}
	fmt.Printf("\nДо сортировки: %v\n", args)
	args = mergeSort(args)
	fmt.Printf("После сортировки: %v\n", args)

	// Быстрая сортировка
	args = []int{21, 1, 22, 10, -5}
	fmt.Printf("\nДо сортировки: %v\n", args)
	args = quickSort(args)
	fmt.Printf("После сортировки: %v\n", args)
}

// Сортировка пузырьком
func bubbleSort(args []int) {
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
}

// Сортировка вставками
func insertionSort(args []int) {
	for i := 1; i < len(args); i++ {
		key := args[i]
		j := i - 1

		for j >= 0 && args[j] > key {
			args[j+1] = args[j]
			j = j - 1
		}
		args[j+1] = key
	}
}

// Сортировка слиянием
func mergeSort(args []int) []int {
	if len(args) <= 1 {
		return args
	}
	middle := len(args) / 2
	left := mergeSort(args[:middle])
	right := mergeSort(args[middle:])

	return merge(left, right)
}

// Сортировка слиянием
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}

		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

// Быстрая сортировка
func quickSort(args []int) []int {
	if len(args) < 2 {
		return args
	}

	left, right := 0, len(args)-1

	pivot := len(args) / 2

	args[pivot], args[right] = args[right], args[pivot]

	for i := range args {
		if args[i] < args[right] {
			args[i], args[left] = args[left], args[i]
			left++
		}
	}

	args[left], args[right] = args[right], args[left]

	quickSort(args[:left])
	quickSort(args[left+1:])

	return args
}
