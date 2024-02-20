package main

func main() {
	// this type of declaration in only valid inside a function
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
