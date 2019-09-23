package main

import "fmt"

func main() {
	x := "Borsh"
	y := "Rabbit"
	z := "USA"

	text := z

	if text == x {
		fmt.Printf("%s equals %s", text, x)
	} else if text == y {
		fmt.Printf("%s equals %s", text, y)
	} else {
		fmt.Printf("%s does not equal anything", text)
	}
}
