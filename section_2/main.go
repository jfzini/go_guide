package main

// var card string = "Ace of Spades"

// func main() {
// 	// this type of declaration in only valid inside a function
// 	cards := []interface{}{"Ace of Diamonds", newCard(), 15, 29}
// 	cards = append(cards, "Six of Spades")
// 	for _, card := range cards {
// 		switch v := card.(type) {
// 		case string:
// 			fmt.Println("String: ", v)
// 		case int:
// 			fmt.Println("Int: ", v)
// 		default:
// 			fmt.Println("Unknown type")
// 		}
// 	}
// 	fmt.Println(cards)
// }

func main() {
	cards := deck{"Ace of Diamonds", newCard(), "Six of Spades"}
	cards = append(cards, "Seven of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
