package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "I'm", "Charsi"}
	// This modifies the value in the slice without using pointers.
	// Why does code function differently with a slice than a struct?
	// Arrays:
	// - Primitive data structure
	// - Can't be resized
	// - Rarely used directly
	// Slices (a fancy array):
	// - Can grow and shrink
	// - Used 99% of the time for lists of elements

	// reference types: (no need to worry about pointers)
	// - slice
	// - maps
	// - channels
	// - pointers
	// - functions

	// value types: (use pointers to change these things in a function)
	// - int
	// - float
	// - string
	// - bool
	// - struct

	// remember, go is a pass-by-value language.

	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func updateSlice(s []string)  {
	s[0] = "Bye"
}