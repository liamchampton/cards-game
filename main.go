package main

// create and manipulate a deck of cards

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	cards := newDeckFromFile("my-cards.txt")
	cards.print()

}
