package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Creating a routine and waitgroup
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Routines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Routines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i += 1 {
		fmt.Println("foo:", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i += 1 {
		fmt.Println("bar:", i)
	}
}
