package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	x := sum("James", xi...)
	fmt.Println("Main total: ", x)
	x = sum("James") //no arguments, but works
	fmt.Println("Main total: ", x)
}

func sum(s string, x ...int) int {
	fmt.Println(s, x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, val := range x {
		sum += val
		fmt.Println("For item in index position, ", i, " we are now adding, ", val, " to the total which is now, ", sum)
	}
	fmt.Println("Final total: ", sum)
	return sum
}
