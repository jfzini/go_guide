package main

func main() {
	// this type of declaration in only valid inside a function
	cards := newDeck()
	hand, remainingDeck := cards.deal(5)
	cards = remainingDeck
	hand.saveToFile("hand.txt")
	cards.saveToFile("deck.txt")
}
