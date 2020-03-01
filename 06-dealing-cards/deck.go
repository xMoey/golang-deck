package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// `cards[0:3]` = range of cards, start at 0 including zero, up to and not including 3.
	// `cards[:3]` = go will infer from the begging up to but not including 3.
	// `cards[2:]` = from index of 2 to the end of the slice.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// functions in go can return multiple values.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
