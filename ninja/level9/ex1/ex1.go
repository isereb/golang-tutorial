package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go doSomething(10)
	go doSomething(10)

	wg.Wait()
}

func doSomething(a int) {
	for i := a; i > 0; i-- {
		fmt.Print(i, " ")
	}
	wg.Done()
}
