package main

import "testing"

func TestNewDeck(t *testing.T) {
	// create new deck
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
