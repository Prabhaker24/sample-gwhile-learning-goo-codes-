package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactinfo
}

type contactinfo struct {
	email string
	zip   int
}

func main() {
	alex := person{
		firstName: "prabhaker",
		lastName:  "saxena",
		contactinfo: contactinfo{
			email: "abc@gmail.com",
			zip:   4}}

	/*	var alex person
		fmt.Println(alex)
		fmt.Printf("%+v", alex)
		var pravhu struct*/
	//	fmt.Println(alex)

	alex.update("prabhu")
	alex.details()
}

func (p *person) update(newname string) {
	p.firstName = newname
	//or
	//(*p).firstName=newname
}

func (p person) details() {
	fmt.Printf("%+v", p)
}
