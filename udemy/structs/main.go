package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	p.firstname = newName
}

func main() {
	// var p1 person
	// p1.firstname = "Agustin"
	// p1.lastname = "Lavarello"
	// p1.inf.email = "agustinlare@gmail.com"
	// p1.inf.zipCode = 1640

	p1 := person{
		firstname: "Agustin Miguel",
		lastname:  "Lavarello",
		contactInfo: contactInfo{
			email:   "agustinlare@gmail.com",
			zipCode: 1640,
		},
	}

	p1.updateName("Agustin")
	p1.printPerson()
}
