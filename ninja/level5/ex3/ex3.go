package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Blue",
		},
		fourWheel: false,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println("---------")

	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t.fourWheel)

	fmt.Println("---------")

	fmt.Println(s.doors)
	fmt.Println(s.color)
	fmt.Println(s.luxury)
}
