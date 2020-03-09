package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool // license to kill
}

func main() {
	// Structs - Composite data type
	p1 := person{
		first: "Bobby",
		last:  "Wragger",
		age:   37,
	}

	// Using other structs within a struct
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(p1, sa1)

	// Anonymous Structs
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Ryan",
		last:  "Griff",
		age:   20,
	}
	fmt.Println(p2)

}
