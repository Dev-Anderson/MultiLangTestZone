package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d receibed %d\n", workerId, x)
	}
}

func main() {
	ch := make(chan int)
	qtdWorkers := 10000

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 1000000; i++ {
		ch <- i
	}
}
