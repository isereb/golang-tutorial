package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Operating Arch:", runtime.GOARCH)
	fmt.Println("Amount of CPUs:", runtime.NumCPU())
	fmt.Println("Amount of routines:", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	wg.Add(1)
	go bar()

	wg.Wait()
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Operating Arch:", runtime.GOARCH)
	fmt.Println("Amount of CPUs:", runtime.NumCPU())
	fmt.Println("Amount of routines:", runtime.NumGoroutine())

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("FOO: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("BAR: ", i)
	}
	wg.Done()

}
