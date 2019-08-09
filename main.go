package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	bob := person{firstName: "Bob", lastName: "Anderson"}
	var bode person

	jim := person{
		firstName: "Jim",
		lastName:  "Partson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	bode.firstName = "Bode"
	bode.lastName = "Fals"

	fmt.Println(bob)

	fmt.Println(bode)
	fmt.Printf("%+v\n", bode)

	jim.updateName("Jimmy")
	jim.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
