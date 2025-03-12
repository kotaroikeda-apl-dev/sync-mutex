package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func setup() {
	fmt.Println("初期化処理実行")
}

func worker(id int) {
	once.Do(setup) // 一度だけ実行
	fmt.Printf("Worker %d started\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go worker(i)
	}

	fmt.Scanln() // 終了を防ぐ
}