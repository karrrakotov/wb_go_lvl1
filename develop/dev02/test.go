package main

import (
	"fmt"
	"sync"
)

/*

=== Задача ===

Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

*/

func main() {
	args := []int{2, 4, 6, 8, 10}

	// Решение (1 способ)
	c := make(chan int)

	for _, val := range args {
		go squares(val, c)
	}

	for range args {
		fmt.Printf("Изначальное число: %v, квадрат этого числа: %v\n", <-c, <-c)
	}

	// Решение (2 способ)
	var wg sync.WaitGroup

	for _, val := range args {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Printf("Изначальное число: %v, квадрат этого числа: %v\n", val, val*val)
		}(val)
	}

	wg.Wait()

}

// Для решения №1
func squares(number int, c chan int) {
	c <- number
	c <- number * number
}
