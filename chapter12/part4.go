package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Mutex condition
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Routines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i += 1 {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v += 1
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Routines:", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
