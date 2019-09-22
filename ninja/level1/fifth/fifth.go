package main

import "fmt"

type myint int

var x myint
var y int

func main() {
	x = 42
	fmt.Printf("[X] Value: %v \n[X] Type: %T \n", x, x)

	y = int(x)
	fmt.Printf("[Y] Value: %v \n[Y] Type: %T \n", y, y)
}
