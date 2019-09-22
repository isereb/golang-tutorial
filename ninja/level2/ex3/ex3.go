package main

import "fmt"

const (
	_ = iota
	a = 10 * iota
	b = 10 * iota
)

const (
	_     = iota
	c int = 21 * iota
	d int = 21 * iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
