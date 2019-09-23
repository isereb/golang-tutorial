package main

import (
	"fmt"
	"time"
)

func main() {
	birthday := 1998
	year := time.Now().Year()

	for {
		if year == birthday {
			break
		}
		fmt.Println(year)
		year--
	}
}
