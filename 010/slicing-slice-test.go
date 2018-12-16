package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x[1])
	fmt.Println(x[1:]) // slicing a slice
	fmt.Println(x[1:2])
}
