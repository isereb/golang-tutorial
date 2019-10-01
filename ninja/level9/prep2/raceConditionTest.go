package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const amt = 100
	counter := 0

	var wg sync.WaitGroup
	wg.Add(amt)

	var mx sync.Mutex

	for j := 0; j < amt; j++ {
		go func() {
			mx.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
