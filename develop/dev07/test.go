package main

import (
	"fmt"
	"sync"
)

/*

=== Задача ===

Реализовать конкурентную запись данных в map.

*/

type SafeMap struct {
	mx sync.Mutex
	m  map[string]string
}

func (c *SafeMap) Store(key string, value string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func main() {
	sm := SafeMap{
		m: make(map[string]string),
	}

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Я суслик №%v и я выполняю сложную работу\n", i)
			defer wg.Done()
			sm.Store(fmt.Sprintf("key%v", i), fmt.Sprintf("val%v", i))
		}(i)
	}

	wg.Wait()

	for key, val := range sm.m {
		fmt.Printf("%s[%s]\n", key, val)
	}
}
