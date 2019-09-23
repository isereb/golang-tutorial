package main

import "fmt"

func main() {
	var arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)
}
