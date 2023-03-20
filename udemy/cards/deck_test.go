package main

import (
	"os"
	"testing"
)

func TestNewDek(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Unexpected length of deck, got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Unexpected card at first position %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Unexpected card at last position %v", cards[0])
	}
}

func TestSaveTodeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 16 {
		t.Errorf("Unexpected length of deck, got %v", len(loadDeck))
	}
}
