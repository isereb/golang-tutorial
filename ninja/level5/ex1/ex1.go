package main

import "fmt"

type person struct {
	first        string
	last         string
	favoriteCars []string
}

func main() {

	james := person{
		first:        "James",
		last:         "Bond",
		favoriteCars: []string{"Jaguar", "Aston Martin"},
	}

	anna := person{
		first:        "Anna",
		last:         "Washington",
		favoriteCars: []string{"Mercedes", "Jeep"},
	}

	displayPerson(james)
	displayPerson(anna)

}

func displayPerson(p person) {
	fmt.Println(p.first, p.last)
	for _, car := range p.favoriteCars {
		println(" -", car)
	}
}
