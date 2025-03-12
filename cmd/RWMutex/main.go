package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeData struct {
	mu   sync.RWMutex
	data int
}

func (s *SafeData) Read(id int) int {
	s.mu.RLock() // 読み取りロック
	defer s.mu.RUnlock()
	fmt.Printf("Reader %d: read data = %d\n", id, s.data)
	return s.data
}

func (s *SafeData) Write(id int, val int) {
	s.mu.Lock() // 書き込みロック
	s.data = val
	fmt.Printf("Writer %d: updated data = %d\n", id, val)
	s.mu.Unlock()
}

func main() {
	safeData := SafeData{}

	for i := 1; i <= 3; i++ {
		go safeData.Read(i)
	}

	time.Sleep(1 * time.Second)
	go safeData.Write(1, 10)

	time.Sleep(1 * time.Second)
}
