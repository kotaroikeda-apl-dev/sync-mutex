package main

import (
	"fmt"
	"time"
)

var counter = 0

func worker(id int) {
	for i := 0; i < 5; i++ {
		counter++ // データ競合発生
		fmt.Printf("Worker %d: counter = %d\n", id, counter)
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
