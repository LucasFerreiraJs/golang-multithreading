package main

// thread 01
func main() {

	forever := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			print(i)
		}

		forever <- true
		// tem que rodar em outra goroutine
	}()

	// sem deadlock
	<-forever
}
