package main

import "fmt"

//noinspection GoBoolExpressions
func main() {
	a := (42 == 31)
	b := (42 <= 31)
	c := (42 >= 31)
	d := (42 != 31)
	e := (42 < 31)
	f := (42 > 31)

	fmt.Println(a, b, c, d, e, f)
}
