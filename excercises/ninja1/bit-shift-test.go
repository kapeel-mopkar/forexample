package main

import (
	"fmt"
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 2 // bit shifting by 2 places in binary * output : 4		100
															//   16		10000
	fmt.Printf("%d\t\t%b\n", y, y)
}