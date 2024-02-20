package main

import "fmt"

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

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
