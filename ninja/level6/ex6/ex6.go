package main

import "fmt"

func main() {
	fmt.Println("Hello world from MAIN!")
	func() {
		fmt.Println("Hello world from anonymous function!")
	}()
}
