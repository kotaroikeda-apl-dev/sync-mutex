package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false
var cond = sync.NewCond(&sync.Mutex{})

func worker(id int) {
	cond.L.Lock()
	for !ready {
		cond.Wait()
	}
	fmt.Printf("Worker %d started\n", id)
	cond.L.Unlock()
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	time.Sleep(2 * time.Second)
	cond.L.Lock()
	ready = true
	cond.Broadcast() // 全員を起こす
	cond.L.Unlock()

	time.Sleep(1 * time.Second)
}
