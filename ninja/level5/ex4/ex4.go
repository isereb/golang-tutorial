package main

import "fmt"

func main() {
	t := struct {
		info   map[string]string
		grades []int
	}{
		info:   make(map[string]string),
		grades: make([]int, 10, 10),
	}

	t.info["name"] = "James"
	t.info["car"] = "Mercedes"

	t.grades = []int{10, 20, 30, 50, 60, 70, 80, 90, 99}

	fmt.Println(t)

	for k, v := range t.info {
		fmt.Println(k, ":\t", v)
	}

	for _, v := range t.grades {
		fmt.Print(v, " ")
	}

}
