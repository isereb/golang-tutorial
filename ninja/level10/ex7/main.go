package main

import (
	"fmt"
)

const first = 10
const second = 10

func main() {
	c := make(chan int)

	for i := 0; i < second; i++ {
		go appendToChanel(i, &c)
	}

	pullAndPrint(&c)
}

func pullAndPrint(c *chan int) {
	for i := 0; i < first*second; i++ {
		val := <-*c
		fmt.Println("Reading from channel (i=", i, ", v=", val, ")")
	}
}

func appendToChanel(start int, c *chan int) {
	for i := start * first; i < start*first+second; i++ {
		fmt.Println("Appending to channel (v=", i, ")")
		*c <- i
	}
}
