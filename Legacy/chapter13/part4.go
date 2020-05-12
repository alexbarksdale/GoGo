package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	go send(c)

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

// send func
func send(c chan<- int) {
	for i := 0; i < 5; i += 1 {
		c <- i
	}
	close(c)
}
