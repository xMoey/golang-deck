package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Golang",
		contactInfo: contactInfo{
			email:   "alex@lol.com",
			zipCode: 1234,
		},
	}
	// you dont need to specifically pass a pointer to `alex` to the
	// `updateName` function, go is smart enough to automatically use
	// the pointer if the type matches.
	alex.updateName("Moe")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
