package main

import "fmt"

// Be careful when changing the order of fields as there is a hidden constructor
// and if you change the order it could result in different values being assigned
// to different variables unexpectedly.
type person struct {
	firstName string
	lastName  string
	// you dont have to init this as `contact contactInfo`, you can do `contactInfo`
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// you can explicitly assign values to variables in the struct like this
	alex := person{
		firstName: "Alex",
		lastName:  "Golang",
		contactInfo: contactInfo{
			email:   "alex@lol.com",
			zipCode: 1234,
		},
	}
	alex.print()

	// reference pointer:
	alexPointer := &alex
	alexPointer.updateName("Moe")
	// john will override Moe since both alexPointer and alex are the exact same
	// struct in memory.
	alex.updateName("John")

	println(alexPointer)

	// you can print the pointer using:
	println(&alex)

	// we print out a person with the first name John twice:
	alex.print()
	alexPointer.print()
}

// if you dont use a pointer `*`, the update of the name wont work on the "object"
// go is a pass-by-value language. this means that if you update something on a struct,
// it takes it from the address its located to another address with the new updated value.
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
