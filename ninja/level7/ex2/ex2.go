package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func main() {

	pers := person{
		first:  "Ivan",
		second: "Sidorov",
		age:    42,
	}

	fmt.Println(pers)
	changeMe(&pers)
	fmt.Println(pers)

}

func changeMe(person *person) {
	person.first = "I am " + person.first
	person.second = "The " + person.second
	person.age++
}
