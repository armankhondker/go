package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spaces"
}

func testing() {
	var card string = "Ace of Spaces" //don't need to write out string here necessarily, go can infer
	var x int = 10
	y := 22 //go infers the type of int
	card2 := "Ace of Diamonds"
	fmt.Println(card + " " + card2)
	fmt.Println(x + y)
}
