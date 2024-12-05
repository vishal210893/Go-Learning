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

	//info := contactInfo{email: "Alex@gmail.com", zipCode: 560103}
	alex := person{
		firstName: "Alex",
		lastName:  "Kumar",
		contactInfo: contactInfo{
			email:   "Alex@gmail.com",
			zipCode: 560103,
		},
	}
	alex.print()
	fmt.Println()

	var ren person
	ren.firstName = "ren"
	ren.lastName = "Williamson"
	fmt.Printf("%+v", ren)
	fmt.Println()

	ren.updateName("Kane")
	ren.print()
	fmt.Println()

	var test string = "Hello"
	t := &test
	fmt.Println(t)
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
