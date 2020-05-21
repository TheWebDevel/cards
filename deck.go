package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

// A function to get some cards based on the hand size
// This function accepts 2 args which are a
// deck of cards d with type deck & hand size of type int
// This function returns two sepearate slices of deck
// The (deck, deck) denotes that we are returning two
// values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// This function has a receiver d of type deck
// and a return type of string.
// The purpose of this function is to conver the
// deck type value to a string type.
func (d deck) toString() string {
	// convert deck to slice of string
	// join the slice of string with the seperator
	// which results in a single string.
	return strings.Join([]string(d), ",")
}

// This a function to save the deck to a file
// which has a receiver d of type deck and accepts
// an argument named filename of type string
// It also can return an error if occurs
func (d deck) saveToFile(filename string) error {
	// We then write the cards to a file
	// This function accepts 3 parameters, namely:
	// 1. filename
	// 2. byte slice (our cards)
	// 3. permissions (who has access and who dont)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// This function creates a deck from a file
// It does not require a receiver because we
// don't have a deck yet which is obvious.
func newDeckFromFile(filename string) deck {
	// This function accepts the file name as a parameter
	//  and do a multiple return the required byte slice & error if any.
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// There are two possible options here,
		// 1. Log the error and return a call to newDeck()
		// 2. Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		// Exit the program entirely
		os.Exit(1)
	}

	// Convert the byte slice into string
	// Convert the string slice of type string []string
	// Then conver the slice of type string to a value of type deck
	return deck(strings.Split(string(bs), ","))
}
