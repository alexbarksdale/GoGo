package main

import (
	"fmt"
)

func main() {
	// NESTED LOOPS
	for i := 0; i <= 5; i += 1 {
		fmt.Printf("The outer loop %d\n", i)
		for j := 0; j <= 3; j += 1 {
			fmt.Printf("\t\tThe inner loop %d\n", j)

		}
	}

	// As long as a bool condition is true a ForStmt will run
	// This is kinda like a while statement in Go
	a := 0
	b := 5

	for a < b {
		fmt.Println("A is < B")
		a += 1
	}

	// Another way to use ForStmt
	x := 0
	for {
		x += 1
		if x > 20 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("done.")

	for i := 33; i <= 122; i += 1 {
		// Binary Hex Unicode
		fmt.Printf("%v\t %#x\t%#U\n", i, i, i)
	}

	// There is no default fall through
	switch {
	case false:
		fmt.Println("No print")
	case (2 == 4):
		fmt.Println("No print")
	case (4 == 4):
		fmt.Println("Print")
		fallthrough
		// Fall through example. Without it, it'll just break
	case (5 == 5):
		fmt.Println("Fall through")
	default:
		fmt.Println("This is the default")
	}
}
