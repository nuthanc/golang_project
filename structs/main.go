package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "Carrie", contactInfo{"a.a@b.com", 562114}}

	alex := person{
		firstName: "Alex", 
		lastName: "Carrie", 
		// contact: contactInfo{
		// 	email: "a.a@b.com", 
		// 	zipCode: 563114,
		// },
		contactInfo: contactInfo{
			email: "a.a@b.com", 
			zipCode: 563114,
		},
	}
	// alexPointer := &alex
	// alexPointer.updateName("Sasuke")
	alex.updateName("Naruto")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
  // But even the below works!
	// pointerToPerson.firstName = newFirstName
}