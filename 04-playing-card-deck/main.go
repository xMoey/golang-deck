// to run this project, `go run main.go deck.go`
package main

func main() {
	// in `deck.go`, we've created a new data struct called deck
	cards := deck{"Ace of Hearts", newCard()}
	cards = append(cards, "Six of Diamond")

	cards.print()
}

func newCard() string {
	return "Five of Clubs"
}