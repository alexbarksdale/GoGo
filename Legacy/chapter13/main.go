package main

import (
	"fmt"
)

func main() {
	// // Channels BLOCK!
	// c := make(chan int)
	// c <- 42 // Gets blocked
	// fmt.Println(<-c)

	// Unblocked Version 1
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	// Unblocked Version 2
	d := make(chan int, 1) // Buffer channel - allows certains channels to sit in it
	d <- 44
	fmt.Println(<-d)

	// Unsuccessful Buffer
	// e := make(chan int, 1)
	// e <- 44
	// e <- 46 // This chan can only hold 1 buffer. Blocked until it's taken off
	// fmt.Println(<-e)
}
