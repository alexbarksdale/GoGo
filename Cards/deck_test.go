package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	len_d := len(d)

	if len_d != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len_d)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but go %v", d[0])
	}

	if d[len_d-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len_d-1])
	}
}

func TestSaveToDeckAndNewDeckFromFrile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	len_loadedDeck := len(loadedDeck)
	if len_loadedDeck != 16 {
		t.Errorf("Expected 16 cards in deck, but go %v", len_loadedDeck)
	}
	os.Remove("_decktesting")

}
