package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	foo()
	bar("Bob")
	s1 := woo("Ricky")
	fmt.Println(s1)
	x, y := too("Gabe ", "Hoski")
	fmt.Println(x, y)
	choo(xi...)
	defer big()
	bendo()
}

// Template: func (r receiver) identifier(params) (returns) {...}
func foo() {
	fmt.Println("Foo")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

// Multiple return
func too(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, " says hello")
	b := true
	return a, b

}

// Variadic parameters
func choo(x ...int) int { // unlimited amount of ints
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding, ", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

// Defer
func big() {
	fmt.Println("big")

}

func bendo() {
	fmt.Println("bendo")
	ha
}
