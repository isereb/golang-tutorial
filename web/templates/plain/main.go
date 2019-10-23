package main

import (
	"fmt"
	"os"
	"text/template"
)

var aTemplate = `
Hi. My name is {{.Name}}
I am  {{.Age}} years old.
`

type person struct {
	Name string
	Age  int
}

func main() {

	// Parse template
	temp, err := template.New("temp").Parse(aTemplate)
	if err != nil {
		fmt.Print("Failed to parse a template: ", err)
	}

	// Create a person
	p := person{
		Name: "Gregory",
		Age:  22,
	}

	// Execute template
	err = temp.Execute(os.Stdout, p)
	if err != nil {
		fmt.Print("Failed to execute a template: ", err)
	}

	// Create another person
	another := person{
		Name: "Alice",
		Age:  19,
	}

	// Execute template with another person
	err = temp.Execute(os.Stdout, another)
	if err != nil {
		fmt.Print("Failed to execute a template: ", err)
	}
}
