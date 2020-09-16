package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}

	// c := make(chan string)
	// go count("sheep", c)

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// for {
	// 	msg, open := <-c

	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
}

// WAIT GROUPS

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		count("sheep")
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }
