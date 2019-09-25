package main

import "fmt"

func main() {
	defer fmt.Println("before dinner")
	fmt.Print("No candies ")
}
