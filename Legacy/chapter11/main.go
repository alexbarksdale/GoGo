package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	// Marshal
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Bob",
		Last:  "Greg",
		Age:   25,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// Unmarshal
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Bob","Last":"Greg","Age":25}]`
	biteSlice := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", biteSlice)

	var people2 []person // Other way: people := []person{}

	err = json.Unmarshal(biteSlice, &people2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The data:", people2)
	for i, v := range people {
		fmt.Println("\nPerson #", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
