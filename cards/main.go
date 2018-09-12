package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "five of diamonds"
}
