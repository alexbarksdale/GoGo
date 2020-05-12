package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Template: func (r receiver) identifier(params) (returns) {...}
// func (s secretAgent) funName() {} - the receiver attaches the function to secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the person speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	// Any type with the method speak is also type Human
	speak()
}

func bar(h human) { // Takes in type human
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(p1)

	// Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
