package main

import "fmt"

func main() {
	c := make(chan int)
	go appendToChanel(&c)
	pullAndPrint(&c)
}

func pullAndPrint(c *chan int) {
	for v := range *c {
		fmt.Println(v)
	}
}

func appendToChanel(c *chan int) {
	for i := 0; i < 100; i++ {
		*c <- i
	}
	close(*c)
}
