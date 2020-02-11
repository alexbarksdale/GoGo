package main

import "fmt"

// Create a new type of 'deck' (a new slice of strings)
type deck []string

// Any variable of type DECK gets access to the printDeck() method
func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
