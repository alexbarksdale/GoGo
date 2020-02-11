package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // append returns a new array

	cards.printDeck()
}

func newCard() string {
	return "Five of Diamonds"
}
