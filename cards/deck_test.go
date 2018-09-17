package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("expected 16 cards, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("expected Ace of Spades, but got ", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Error("expected last card of Fource of Diamonds, but got : ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("expected 16 cards, but got ", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
