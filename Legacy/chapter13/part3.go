package main

import (
	"fmt"
)

// Using select to pull values
func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)
	fmt.Println("Exiting....")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i += 1 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From even", v)
		case v := <-o:
			fmt.Println("From odd", v)
		case v := <-q:
			fmt.Println("From quit", v)
			return
		}
	}
}
