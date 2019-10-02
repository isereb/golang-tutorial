package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// c := make(chan int, 1) // this also works without anon func

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
