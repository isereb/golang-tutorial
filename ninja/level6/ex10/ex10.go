package main

import "fmt"

func main() {
	i := iter(10)
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

	it := iter(12)
	fmt.Println(it())
	fmt.Println(it())
	fmt.Println(it())
	fmt.Println(it())
}

func iter(amt int) func() int {
	x := 0
	return func() int {
		x += amt
		return x
	}
}
