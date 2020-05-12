package main

import "fmt"

var y = 42
var a int

type hotdog int

var b hotdog

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Printf("Variable declared with type\n")

	a = 55
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Printf("Custom type\n")

	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Printf("Converting a type\n")
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
