package main

import "fmt"

// describe what a deck is and how it works

// create a new type of 'deck' - slice of strings

type deck []string

// create and return a list of playing cards (a deck)
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// split deck with hand size and return 2 deck's
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
