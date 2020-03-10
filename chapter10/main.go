package main

import "fmt"

func main() {
	// Pointers
	fmt.Println("Pointers:")
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you an address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)  // * derefrences an address Ex: 0xc0000120a0 -> 42
	fmt.Println(*&a) // Gives address (&), then give you the value stored (*)

	*b = 43        // Reassigns the value stored at b to 43
	fmt.Println(a) // a & b are pointing to the same address

	// Pointers in action
	fmt.Println("Pointers in action:")
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x) // Gives the address to the value of x
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43 // Derefrences the value stored at y (x)
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
