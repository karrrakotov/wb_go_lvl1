package main

/*

=== Задача ===

Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

*/

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	jobs := make(chan int)

	n := 5
	for w := 1; w < n; w++ {
		go worker(w, jobs)
	}

	go func() {
		for i := 1; ; i++ {
			jobs <- i
		}
	}()

	// Отслеживание сигналов операционной системы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	<-signalChan

	// Завершение работы воркеров
	close(jobs)
	fmt.Println("Программа завершена")
}

func worker(id int, jobs chan int) {
	for i := range jobs {
		fmt.Println("Worker", id, "processed job", i)
	}
}
