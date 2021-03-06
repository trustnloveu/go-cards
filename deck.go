package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

//! Receiver
func (d deck) print() {
	for i, card := range d { // 'd' can be replaced with 'this' or 'self'
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 : Anyone can read and write the file
}

func (d deck) shuffle() {
	//* Seed Value, which is based on Unix time in a type of int64
	source := rand.NewSource(time.Now().UnixNano()) // time pkg > Now > UnixNano
	r := rand.New(source)

	for i := range d {
		//* Generate random number
		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)

		//* Swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

//! Not Receiver
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // bs = byteSlice, err = If nothing went wrong, it will have a value of 'nil'

	if err != nil {
		// Option #1 - log the error and return a call to newDeck()

		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error ::: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") // strings.Split > return []string
	return deck(s)
}
