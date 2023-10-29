package main

import (
	"fmt"
	"sync"
)

/*

=== Задача ===

Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+4^2+6^2...) с использованием конкурентных вычислений.

*/

func main() {
	args := []int{2, 4, 6, 8, 10}

	sum := sum(args)
	fmt.Printf("Входные данные: %v, результат суммы: %v\n", args, sum)
}

// TODO - 1 Способ
func sum(args []int) (sum int) {
	var wg sync.WaitGroup

	for _, val := range args {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sum += val * val
		}(val)
	}

	wg.Wait()

	return sum
}
