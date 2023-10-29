package main

import (
	"fmt"
	"sync"
)

/*

=== Задача ===

Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

*/

func main() {
	args := [5]int{1, 5, 12, 6, 8}
	firstChannel := make(chan int)
	secondChannel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			x, open := <-firstChannel

			if !open {
				close(secondChannel)
				return
			}

			secondChannel <- x * x
		}
	}()

	go func() {
		defer wg.Done()
		for {
			x, open := <-secondChannel

			if open {
				fmt.Printf("Полученные данные с второго канала: %v\n", x)
			} else {
				return
			}
		}
	}()

	for _, val := range args {
		firstChannel <- val
	}

	close(firstChannel) // Закрываем первый канал после отправки всех данных
	wg.Wait()           // Ожидаем завершения обеих горутин
}
