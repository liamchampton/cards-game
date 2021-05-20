package main

// create and manipulate a deck of cards

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	cards := newDeck()

	hand, remmainingCards := deal(cards, 5)

	hand.print()
	remmainingCards.print()
}
