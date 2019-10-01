package main

import (
	"fmt"
	"sync"
)

func main() {

	amt := 100
	counter := 0

	var wg sync.WaitGroup
	wg.Add(amt)
	for j := 0; j < amt; j++ {
		fmt.Println("START: Iteration: ", j, "Value:", counter)
		go func() {
			v := counter
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("END: Iteration: ", j, "Value:", counter)
	}
	wg.Wait()
}
