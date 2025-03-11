package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu  sync.Mutex
	val int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
	fmt.Printf("Counter: %d\n", c.val)
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func main() {
	var wg sync.WaitGroup
	counter := SafeCounter{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter.Value())
}
