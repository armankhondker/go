package main

func main() {
	// var card string = "Aces of Spades"
	// fmt.Println(card)
	// card1 := "Aces of Diamonds"
	// fmt.Println(card1)
	// card1 = "Five of Diamonds"
	// fmt.Println(card1)
	// fmt.Println(helper())
	// tester := helper()
	// fmt.Println(tester)
	// cards := []string{"Arman", "Ganiyu", "Punit"}
	// cards = append(cards, "Test")
	// fmt.Println(cards)
	// mycards := deck{"Kai", "Lindsey", "Jianfeng"}
	// mycards.print()
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
