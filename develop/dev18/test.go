package main

/*

=== Задача ===

Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.

*/

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterWithAtomic struct {
	counter uint64
}

type CounterWithMutex struct {
	rw      sync.RWMutex
	counter uint64
}

// Вижу несколько решений:
// 1. (Предпочтительный вариант) - использовать Atomic для синхронизации, т.к. они были созданы для этого
// 2. Использовать обычные Mutex или RWMutex (я выберу их)
func (c *CounterWithAtomic) Increment() {
	atomic.AddUint64(&c.counter, 1)
}

func (c *CounterWithAtomic) Value() uint64 {
	return atomic.LoadUint64(&c.counter)
}

func (c *CounterWithMutex) Increment() {
	c.rw.Lock()
	c.counter++
	c.rw.Unlock()
}

func (c *CounterWithMutex) Value() uint64 {
	c.rw.RLock()
	defer c.rw.Unlock()
	return c.counter
}

func main() {
	// Реализация с помощью Atomic
	var counterWithAtomic CounterWithAtomic = CounterWithAtomic{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counterWithAtomic.Increment()
		}()
	}

	wg.Wait()

	fmt.Printf("Счетчик равен: %v\n", counterWithAtomic.counter)

	// Реализация с помощью RWMutex
	var counterWithMutex CounterWithMutex = CounterWithMutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counterWithMutex.Increment()
		}()
	}

	wg.Wait()

	fmt.Printf("Счетчик равен: %v\n", counterWithMutex.counter)
}
