package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%d -> %#U\n", j+1, i)
		}
		fmt.Println()
	}
}
