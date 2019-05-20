package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))

	}

	if d[0] != "Ace of Speads" {
		t.Errorf("Expected first card of Ace of speade, but got %v", d[0])

	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of four of Clubs, but got %v", d[len(d)-1])

	}
}

func TestSaveTodeck(t * testing.T) {
	os.Remove("_desktesting")

	deck := newDeck()
	deck.toSaveFile("_desktesting")
	

	loadedDeck := newDeckFile("_desktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 card of deck, but got %v ", len(loadedDeck));
	}
	os.Remove("_desktesting")

}
