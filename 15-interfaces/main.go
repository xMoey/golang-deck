package main

import "fmt"

// types:
// concrete types:
// - map, struct, int, string, englishBot
// interface type:
// - bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// once you implement one function from the interface, its assumed that "class"
// now implements that interface so you need to implement all the functions
// defined in the interface.
func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}