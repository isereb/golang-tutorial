package main

import "fmt"

func main() {
	c := make(chan int)
	go appendToChanel(&c)
	pullAndPrint(&c)
}

func pullAndPrint(c *chan int) {
	for v := range *c {
		fmt.Println("Reading from channel (i=", v, ")")
	}
}

func appendToChanel(c *chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println("Appending to channel (i=", i, ")")
		*c <- i
	}
	close(*c)
}
