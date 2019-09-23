package main

import (
	"fmt"
	"time"
)

func main() {
	birthday := 1998
	year := time.Now().Year()
	for birthday <= year {
		fmt.Println(birthday)
		birthday++
	}
}
