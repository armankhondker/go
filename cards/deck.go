package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string //deck kind of extends string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {

}

func (d deck) deal() {

}

func (d deck) saveToDeal() {

}

func (d deck) newDeckFromFile() {

}
