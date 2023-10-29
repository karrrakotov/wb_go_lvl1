package main

/*

=== Задача ===

Реализовать все возможные способы остановки выполнения горутины.

*/

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Остановка по истечении таймера
	go func() {
		fmt.Println("Горутина: Ожидание в течение 3 секунд...")
		time.Sleep(3 * time.Second)
		fmt.Println("Горутина: Завершение по таймауту")
	}()

	// 2. Остановка по каналу
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Горутина: Завершение по сигналу из канала")
				return
			default:
				fmt.Println("Горутина: Выполнение работы...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 3. Остановка по сигналу операционной системы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Горутина: Завершение по сигналу операционной системы:", sig)
	}()

	// Для демонстрации, программа будет выполняться в течение 10 секунд
	time.Sleep(10 * time.Second)

	// Остановка по каналу
	stop <- true

	// Остановка программы при получении сигнала
	fmt.Println("Главная программа: Завершение программы.")
}
