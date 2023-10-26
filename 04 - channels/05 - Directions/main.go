package main

import "fmt"

// * direita, atribui informação <-
func recebe(nome string, helloch chan<- string) {
	helloch <- nome
}

// * <- esquerda, esvazia canal
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	helloch := make(chan string)
	go recebe("hello world", helloch)
	ler(helloch)
}
