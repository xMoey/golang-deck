package main

import "fmt"

// go is not an oop language and does not have class or anything like that
// every time you see `type xxx xxx` think of it as the "class" with the functions
// implemented below it.

// Create a new type of 'deck' which is a slice of strings
// you can think of this as our new `deck` type extends `[]string`
// aka a "slice of strings"
type deck []string

// if you add a name before the function name e.g. `(d deck)`
// you've added a receiver.
// `d` is the instance of the deck we are working with, its just a variable btw.
// think of `d` as a `this` or `self`.  Convention is 1-3 letter abbreviation
// of obj name.
// `deck` means we've attached it to every var of type deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
