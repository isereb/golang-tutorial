package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Hello world from anonymous function!")
	}

	fmt.Println("Hello world from MAIN!")

	x()
}
