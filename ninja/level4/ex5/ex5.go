package main

import "fmt"

func main() {
	//		   0   1   2   3   4   5   6   7   8   9
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:3], x[6:]...)
	fmt.Println(y)

	//		   0   1   2   3   4   5   6   7   8   9
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Delete 46 at index 4
	a[4] = a[len(a)-1]
	b := a[:len(a)-1]
	fmt.Println(b)

}
