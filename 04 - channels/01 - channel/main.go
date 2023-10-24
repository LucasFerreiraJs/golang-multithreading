package main

import "fmt"

// * Thread 1
func main() {

	canal := make(chan string) // * vazio

	// * thread 2
	go func() {
		canal <- "teste channel" // * cheio
	}()

	msg := <-canal // * esvazia canal

	fmt.Println(msg)

}
