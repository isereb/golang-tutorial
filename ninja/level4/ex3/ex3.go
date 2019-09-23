package main

import "fmt"

func main() {
	//			 0	 1   2   3   4   5   6   7   8   9
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := arr[:5]
	b := arr[5:]
	c := arr[2:7]
	d := arr[1:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
