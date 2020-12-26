package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	d.TestDeck(t)

}

func TestSavetoDeckfromLoadDeck(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := loadDeck("_decktesting")

	loadedDeck.TestDeck(t)

	os.Remove("_decktesting")

}

func (d deck) TestDeck(t *testing.T) {
	if len(d) != 52 {
		t.Errorf("Expected 52 cards, got %v", len(d))
		os.Remove("_decktesting")
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Wrong first card. Got %v, expected 'Ace of Spades'", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Wrong first card. Got %v, expected 'King of Clubs'", d[len(d)-1])
	}
}
