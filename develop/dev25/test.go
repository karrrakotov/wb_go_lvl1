package main

/*

=== Задача ===

Реализовать собственную функцию sleep.

*/

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начало программы")
	sleep(10 * time.Second)
	fmt.Println("Конец программы")
}

// Своя реализация sleep
func sleep(d time.Duration) {
	<-time.After(d)
}
