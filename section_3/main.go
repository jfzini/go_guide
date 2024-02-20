package main

func main() {
	// this type of declaration in only valid inside a function
	cards := newDeck()
	hand, remainingDeck := cards.deal(5)
	hand.print()
	cards = remainingDeck
	cards.print()
}
