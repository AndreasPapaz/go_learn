package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// cards := deck{newCard(), "Ace of Diamonds"}
	// cards = append(cards, "Six of Spades")
	// cards := newDeck()

	// cards.print()
	// hand, remCards := deal(cards, 5)

	// hand.print()
	// remCards.print()

	// cards.saveToFile("my_cards")

	cards := readFromFile("my_cards")
	cards.print()
}

func newCard() string {
	return "five of diamonds"
}
