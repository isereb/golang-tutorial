package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 1234
}

func bar() (int, string) {
	return 700, "James"
}
