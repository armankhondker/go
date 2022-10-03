package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string //deck kind of extends string

func newDeck() deck {
	cards := deck{} //make an empty array
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile() {
}

func (d deck) toString() string { //converts deck which is an array of strings, into one large string
	return strings.Join([]string(d), ",")
}

func (d deck) newDeckFromFile() {

}
