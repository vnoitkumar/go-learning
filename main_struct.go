package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string
	contact   *contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Kumar",
		contact: &contactInfo{
			email:   "jimkumar@gmail.com",
			pincode: 635701,
		},
	}
	jim.print()
	jim.updateName("Vnoit")
	jim.print()
}
