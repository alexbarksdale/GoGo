package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// WAYS OF INIT A STRUCT TYPE
	// ===================================

	// 1. This relies heavily on the order of the struct.
	// rob := person{"Rob", "Ery"}

	// 2. Defines the fields of the struct
	// rob := person{firstName: "Rob", lastName: "Ery"}
	// fmt.Println(alex)

	// 3. Go assigns a 0 value when creating var of type struct like this
	//    0 value examples: strings = "" | int = 0 | bool = false
	// 	  Ex: firstName = "" && lastName = ""
	// var rob person
	// fmt.Println(rob)
	// fmt.Printf("%+v", rob) // %+v = print all the fields and their values

	// rob.firstName = "Rob"
	// rob.lastName = "Ery"
	// fmt.Printf("\n%+v", rob)

	jim := person{
		firstName: "Jim",
		lastName:  "Emy",
		contactInfo: contactInfo{
			email:   "jim@test.com",
			zipCode: 76321,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	jim.updateName("jimmy")
	jim.info()
}

func (p person) info() {
	fmt.Printf("\n%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
