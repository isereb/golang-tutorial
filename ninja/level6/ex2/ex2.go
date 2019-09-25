package main

import "fmt"

func main() {
	numbers1 := []int{1, 3, 5, 1, 8, 2, 5, 1, 5, 8, 87, 65}
	sum1 := foo(numbers1...)
	fmt.Println("[1] Sum of numbers: ", sum1)

	numbers2 := []int{7247, 12341, 87714, 174714, 4241}
	sum2 := foo(numbers2...)
	fmt.Println("[2] Sum of numbers: ", sum2)
}

func foo(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func bar(ii []int) int {
	return foo(ii...)
}
