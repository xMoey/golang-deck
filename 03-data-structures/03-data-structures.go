package main

import "fmt"

func main() {
	// Slices vs Arrays.
	// in an array, elements can be of whatever type.
	// you can mix an int with a string. slices are different,
	// every element must be of the same type.
	cards := []string{"Ace of Hearts", newCard()}
	cards = append(cards, "Six of Diamond")

	// `i` = index of element.
	// `card` = current card we're iterating over.
	// `range cards` = range is a keyword to iterate over
	//  every element in the slice.
	// `:=` is used in for loops because each time we go to a new
	//  element we are abandoning the last element and need a new one.
	for i, card := range cards {
		fmt.Println(i , card)
	}
}

func newCard() string {
	return "Five of Clubs"
}