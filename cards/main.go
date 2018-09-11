package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "five of diamonds"
}
