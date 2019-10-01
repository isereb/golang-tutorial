package main

import "fmt"

func main() {

	amt := 100
	counter := 0

	for j := 0; j < amt; j++ {
		fmt.Println("START: Iteration: ", j, "Value:", counter)
		go func() {
			v := counter
			v++
			counter = v
		}()
		fmt.Println("END: Iteration: ", j, "Value:", counter)
	}
}
