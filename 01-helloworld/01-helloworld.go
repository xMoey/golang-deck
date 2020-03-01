// package main = executable file.
package main
// you can also have packages as helpers,
// just name them as package <anything else> = dependency package, reusable.

// good reference for go standard packages: https://golang.org/pkg/flag/
import "fmt"

// every single `package main` must have a `func main()` in the script to work.
// func = tells we are about to declare a function
// main = sets the function name
// () = list of arguments
// {} = function body
func main() {
	message := greetMe("world")
	fmt.Println(message)
}

func greetMe(name string) string {
	return "hello " + name + "!"
}