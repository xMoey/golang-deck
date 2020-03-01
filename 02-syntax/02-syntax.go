// functions in the same package can freely use other functions in other files
// that are also in the same package.
package main

import "fmt"

// global variables cannot be assigned a value during initialisation.
var deckSize int

func main() {
	// you can assign a value to global variables in a function.
	deckSize = 50
	fmt.Print(deckSize)

	// `:=` expanded is `var card string = "Ace"`.
	// go is smart, it can infer types from the value.
	card := newCard()
	fmt.Print(card)
}

// breakdown of a function:
// func <name> <return type> {
func newCard() string {
	return "Five of Diamonds"
}

func estimatePi() float64 {
	return 3.14
}