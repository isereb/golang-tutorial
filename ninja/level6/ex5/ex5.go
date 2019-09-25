package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (c square) area() float64 {
	return c.length * c.width
}

func main() {
	sq := square{
		length: 12.5325,
		width:  18.2421,
	}

	ci := circle{
		radius: 12.1241,
	}

	fmt.Println("Square area:", sq.area())
	fmt.Println("Circle area:", ci.area())
}
