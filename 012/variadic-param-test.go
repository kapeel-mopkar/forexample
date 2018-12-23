package main

import (
	"fmt"
)

func main() {

	x := foo(2, 3, 4, 5, 7, 8, 9)
	fmt.Println("Main total: ", x)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, val := range x {
		sum += val
		fmt.Println("For item in index position, ", i, " we are now adding, ", val, " to the total which is now, ", sum)
	}
	fmt.Println("Final total: ", sum)
	return sum
}
