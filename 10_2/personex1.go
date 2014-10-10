package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1- struct as a value type
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2- struct as a pointer
	var pers2 Person
	pers2.firstName = "Nicole"
	pers2.lastName = "Kidman"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3- struct literal
	pers3 := Person{"Ron", "Howard"}
	upPerson(pers3)
	fmt.Printf("the name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
