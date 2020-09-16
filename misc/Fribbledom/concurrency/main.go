package main

import (
	"fmt"
	"time"
)

func Fibonacci(c chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
		fmt.Println("Pre-gen next number: ", x)
	}
}

func Printer(id int, c chan int) {
	for {
		fmt.Printf("Printer %d got %d\n", id, <-c)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan int)

	go Fibonacci(c)
	go Printer(1, c)
	go Printer(2, c)

	time.Sleep(5 * time.Second)
}
