package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// * thread 1
func main() {
	// * thread 2
	go task("tarefa A")
	// * thread 3
	go task("tarefa B")
	// * thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: task %s is running \n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	// * sai e encerra threads

	// segura a execução
	time.Sleep(15 * time.Second)
}
