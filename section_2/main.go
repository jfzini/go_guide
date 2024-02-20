package main

import "fmt"

// var card string = "Ace of Spades"

func main() {
	// this type of declaration in only valid inside a function
	cards := []interface{}{"Ace of Diamonds", newCard(), 15, 29}
	cards = append(cards, "Six of Spades")
	for _, card := range cards {
		switch v := card.(type) {
		case string:
			fmt.Println("String: ", v)
		case int:
			fmt.Println("Int: ", v)
		default:
			fmt.Println("Unknown type")
		}
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
