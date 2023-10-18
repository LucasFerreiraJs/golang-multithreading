package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, w *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
	}

	w.Done()
}

func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(4)

	go task("A", &waitGroup)
	go task("b", &waitGroup)
	go task("c", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: task %s is running \n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
