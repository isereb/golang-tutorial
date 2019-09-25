package main

import "fmt"

func main() {
	getOne := getFunc()
	fmt.Println(getOne())
}

func getFunc() func() int {
	return func() int {
		return 1
	}
}
