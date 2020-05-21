package main

import "testing"

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
}
