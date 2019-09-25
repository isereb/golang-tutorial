package main

import (
	"fmt"
	"strings"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak(text ...string) string {
	return `"` + strings.Join(text, " ") + "\" - said " + p.first + " " + p.last
}

func main() {
	pers := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(pers.speak("I'm", fmt.Sprint(pers.age), "years old"))
}
