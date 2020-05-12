package main

import (
	"fmt"
)

func main() {
	foo()
	// Anonymous functions
	fmt.Println("\nAnonymous functions:")
	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)

	// Function expresions
	fmt.Println("\nFunction expressions:")
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)

	// Returning a function
	fmt.Println("\nReturning a function:")
	fmt.Println(bar()) // Bar returns a string

	x := voo() // Returns a function
	fmt.Printf("%T", x)
	fmt.Println("\n", x())
	fmt.Println("\n", voo()()) // Other way to do above ^
}

func foo() {
	fmt.Println("foo ran")
}

func bar() string {
	return "Hello world"
}

// Returns a function
func voo() func() int {
	return func() int {
		return 451
	}
}
