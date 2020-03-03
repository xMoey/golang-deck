package main

func main() {
	// maps can be defined:
	// var colors map[string]string
	// colors := make(map[string]string)

	// `[string]` = keys type
	// `string` = value type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#abc123",
	}

	// add element to map:
	colors["white"] = "#ffffff"

	// delete element
	delete(colors, "white")

	printMap(colors)
}

// you can iterate over a map similarly to how we iterate over
// a slice.
func printMap(c map[string]string) {
	for color, hex := range c {
		println(color, hex)
	}
}

//- Maps
//All keys must be the same type
//All values must be of the same type
//Keys are indexed - we can iterate over them
//Use to represent a collection of related properties
//Don't nee dto know all the keys at compaile time
//Reference Type!
//
//- Structs
//Values can be of different type
//Keys don't suport indexing
//You need to know all the different fields at compile time
//Think of these as more like classes in OO programming
//Value Type!