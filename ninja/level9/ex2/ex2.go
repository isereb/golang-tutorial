package main

import "fmt"

type human struct {
	first string
	age   int
}

func speak(p *person) {
	fmt.Println("I'm", p.first)
}

type person human

/* Does not work
func doIt(human human) {
	human.speak()
}
*/

func main() {
	ivan := person{
		first: "Ivan",
		age:   20,
	}

	speak(&ivan) // can do
	// speak(ivan)  // can't do

}
