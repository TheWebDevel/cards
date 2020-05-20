package main

import "fmt"

// Create a new type of deck which
// is a slice of strings
// We tell to Go that there's a new type
// called deck which has the same behaviour
// as a slice of string []string
type deck []string

// Create and return a list of playing cards
// This function will have a return type of deck (our custom type)
// This function does not require a receiver because,
// it'll not be called like cards.newDeck()
func newDeck() deck {
	// Create a new variable called cards of type deck
	cards := deck{}

	// We use string instead of deck because deck is meant
	// to be used with actual playing cards like Ace of Dimonds
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Loop over cardSuits and cardValues
	// and create combinations of eachother
	// Use _ to tell the Go compiler that we
	// don't care about the unused variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Loop through the dech of cards and print
// out the values of every single card
// This is a receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
