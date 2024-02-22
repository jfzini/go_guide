package main

import (
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, value := range cardValues {
		for _, suit := range cardSuits {
			newCard := value + " of " + suit
			cards = append(cards, newCard)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	newHand := d[:handSize]
	remainingDeck := d[handSize:]
	return newHand, remainingDeck
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	sliceBytes := []byte(d.toString())
	return os.WriteFile(filename, sliceBytes, 0666)
}

// func (d deck) print() {
// 	for _, card := range d {
// 		fmt.Println(card)
// 	}
// }
