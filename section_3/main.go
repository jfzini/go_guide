package main

func main() {
	// this type of declaration in only valid inside a function
	cards := newDeck()
	dealResult := deal(5, cards)
	dealResult[0].print()
	dealResult[1].print()
}
