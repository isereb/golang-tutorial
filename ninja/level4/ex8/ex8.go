package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`, `extra element`},
	}

	// fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		for key, value := range v {
			fmt.Println(key, ": ", value)
		}
	}

}
