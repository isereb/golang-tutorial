package main

import "fmt"

type myerr struct {
	info string
}

func (err myerr) Error() string {
	return fmt.Sprintf("There is an error: %v", err.info)
}

func main() {

	myerr := myerr{
		info: "Need more coffee",
	}

	foo(myerr)

}

func foo(e error) {
	fmt.Println("Something went wrong: ", e.(myerr).info)
}
