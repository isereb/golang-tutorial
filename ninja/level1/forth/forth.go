package main

import (
	"fmt"
)

type myint int

var x myint

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
