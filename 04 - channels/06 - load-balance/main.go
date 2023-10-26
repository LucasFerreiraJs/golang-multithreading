package main

import (
	"fmt"
	"time"
)

// * leitura de dados em paralelo através de workers

// * responsável pelo processamento
func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("worker: %d, received: %d\n", workerId, x)
		time.Sleep(time.Second)
	}

}

func main() {
	data := make(chan int)
	qtdWorkers := 10

	// go worker(1, data)
	// go worker(2, data)

	// * inicializa workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
