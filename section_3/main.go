package main

func main() {
	// this type of declaration in only valid inside a function
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
