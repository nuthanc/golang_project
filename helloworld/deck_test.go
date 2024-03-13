package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Error("Expected deck length of 16, but got", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Hearts" {
		t.Errorf("Expected Four of Hearts, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected 16 cards in deck, but got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestShuffle(t *testing.T) {
	cards := newDeck()
	cards.shuffle()
	if cards[0] == "Ace of Clubs" && cards[len(cards) - 1] == "Four of Hearts" {
		t.Error("Not shuffled properly")
	}
}