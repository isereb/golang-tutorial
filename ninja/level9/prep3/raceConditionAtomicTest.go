package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	const amt = 100
	var counter int64

	var wg sync.WaitGroup
	wg.Add(amt)

	for j := 0; j < amt; j++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			//fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
