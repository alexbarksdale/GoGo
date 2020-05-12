package main

import (
	"fmt"
)

func main() {
	// Using channels
	c := make(chan int)

	// send
	go send(c)

	// receive
	receive(c)

	fmt.Println("About to exit")
}

// send func
func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
