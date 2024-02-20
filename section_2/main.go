package main

import "fmt"

// var card string = "Ace of Spades"

func main() {
	// this type of declaration in only valid inside a function
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
