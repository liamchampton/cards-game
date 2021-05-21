package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// create new deck and assign variables
	deck := newDeck()
	firstCard := deck[0]
	lastCard := deck[len(deck)-1]

	// check deck has corrent number of cards - if it doesn't then error
	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	// check first card in deck = ace of spades
	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", firstCard)
	}

	// check last card in deck = four of clubs
	if lastCard != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
