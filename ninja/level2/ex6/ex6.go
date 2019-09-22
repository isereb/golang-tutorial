package main

import "fmt"

const year int = 2019

const (
	_     = iota
	year1 = year + iota
	year2 = year + iota
	year3 = year + iota
	year4 = year + iota
)

func main() {

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
