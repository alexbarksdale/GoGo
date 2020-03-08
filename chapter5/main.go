package main

import (
	"fmt"
)

var x int
var y float64

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

const (
	_ = iota // 0 iota
	// kilobyte = 1024
	kilobyte = 1 << (iota * 10) // 1 iota
	// megabyte = kilobyte * kilobyte
	megabyte = 1 << (iota * 10) // 2 iota
	// gigabyte = kilobyte * megabyte
	gigabyte = 1 << (iota * 10) // 3 iota
)

func main() {
	// Numeric types
	x = 42
	y = 42.3423

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// IOTA - Increments by one automatically
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// Bitshifting
	k := 4
	fmt.Printf("Decimal numbering system: %d\t\tBinary numbering sytem: %b", k, k)

	t := k << 1
	fmt.Printf("\nDecimal numbering system: %d\t\tBinary numbering sytem: %b", t, t)

	fmt.Printf("\nDecimal numbering system: %d\t\tBinary numbering sytem: %b", kilobyte, kilobyte)
	fmt.Printf("\nDecimal numbering system: %d\t\tBinary numbering sytem: %b", megabyte, megabyte)
	fmt.Printf("\nDecimal numbering system: %d\t\tBinary numbering sytem: %b", gigabyte, gigabyte)

}
