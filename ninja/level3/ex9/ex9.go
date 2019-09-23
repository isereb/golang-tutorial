package main

import "fmt"

func main() {

	str := "borsh"
	switch str {
	case "pelmeni":
		fmt.Println("Pelmeniiiiii")
	case "borsh":
		fmt.Println("Borshez")
	case "pivo":
		fmt.Println("Pivchik")
	default:
		fmt.Println("You don't get any food!")
	}
}
