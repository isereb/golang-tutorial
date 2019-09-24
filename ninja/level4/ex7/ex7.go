package main

import "fmt"

func main() {
	jb := []string{"James", "Bond    ", "Shaken, not stirred."}
	mp := []string{"Miss ", "Moneypenny", "Helloooooo, James."}

	ppl := [][]string{jb, mp}

	for i, v := range ppl {
		fmt.Print(i, ": ")
		for _, value := range v {
			fmt.Print(value, "\t")
		}
		fmt.Println()
	}
}
