package main

import (
	"os"
	"testing"
)

// Go will automatically call this function
//  with a parameter t *testing.T Here, the t is our
// test handler. So, we'll be notifying the test handler
// if something goes wrong.
func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	if len(d) != 16 {
		// Notify the test handler
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	// Verify first element of the deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	// Verify last element of the deck
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[0])
	}
}

// A function to test save deck to file
//  and load deck from file
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove files
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(loadedDeck))
	}

	// Remove files
	os.Remove("_decktesting")
}
