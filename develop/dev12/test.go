package main

import (
	"fmt"
	"sync"
)

/*

=== Задача ===

Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

*/

func main() {
	args := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("Собственное множество: %v\n", args)
	args = set(args)
	fmt.Printf("Собственное множество: %v\n", args)
}

func set(args []string) (result []string) {
	mapSet := make(map[string]bool)
	var wg sync.WaitGroup

	for _, val := range args {
		wg.Add(1)

		go func(val string) {
			defer wg.Done()
			_, exist := mapSet[val]
			if !exist {
				mapSet[val] = true
				result = append(result, val)
			}
		}(val)
	}

	wg.Wait()
	return
}
