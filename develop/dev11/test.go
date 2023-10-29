package main

import (
	"fmt"
	"sync"
)

/*

=== Задача ===

Реализовать пересечение двух неупорядоченных множеств.

*/

func main() {
	args1 := []int{1, 9, 5, 3, 3, 2, 1}
	args2 := []int{2, 2, 3, 6, 7, 9, 0, 1}

	result := intersection(args1, args2)
	fmt.Printf("Пересечение двух неупорядоченных множеств: %v\n", result)
}

func intersection(args1, args2 []int) (res []int) {
	interMap := make(map[int]bool)
	var wg sync.WaitGroup

	for _, val := range args1 {
		wg.Add(1)
		go func(val int) {
			interMap[val] = true
			wg.Done()
		}(val)
	}

	wg.Wait()

	for _, val := range args2 {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			if interMap[val] {
				res = append(res, val)
				interMap[val] = false
			}
		}(val)
	}

	wg.Wait()
	return
}
