package main

import "fmt"

func main() {
	some := getSome(func() int {
		return 1824
	}, 123)
	fmt.Println(some)
}

func getSome(getSome func() int, some int) int {
	return getSome() * 10
}
