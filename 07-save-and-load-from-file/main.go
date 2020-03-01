package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("test")
	cards := newDeckFromFile("test")
	cards.print()
}
