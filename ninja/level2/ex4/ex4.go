package main

import "fmt"

func main() {
	x := 12
	fmt.Printf("Decimal: %d,\nBinary: %b,\nHex: %x\n", x, x, x)
	y := x << 1
	fmt.Printf("Decimal: %d,\nBinary: %b,\nHex: %x\n", y, y, y)
}
