package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mu sync.Mutex // Mutex 変数

func worker(id int) {
	for i := 0; i < 5; i++ {
		mu.Lock() // ロック
		counter++
		fmt.Printf("Worker %d: counter = %d\n", id, counter)
		mu.Unlock() // アンロック
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Final Counter:", counter)
}
