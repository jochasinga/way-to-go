package main

import (
	"fmt"
)

type VCard struct {
	name, birthdate, photo   string
	addresses                []*Address
}

type Address struct {
	street, apt, city, country, zipcode string

	
}

func main() {
	contact2 := &VCard{
		"Joe", 
		"12/06/83", 
		"image.jpg",
		[]*Address{
			&Address{ 
				"449 W 125th St",
				"4C", "New York",
				"US", "10027",
			},
			&Address{ 
				"609 W 96th St",
				"4I", "New York",
				"US", "10026",
			},
			&Address{ 
				"405 E 304th St",
				"5C", "New York",
				"US", "10032",
			},
		},
	}
	fmt.Println(contact2)
	fmt.Printf("%s was born on %s and is currently living at %s, %s, %s\n", contact2.name, contact2.birthdate, contact2.addresses[0].street, contact2.addresses[0].apt, contact2.addresses[0].city)
}
