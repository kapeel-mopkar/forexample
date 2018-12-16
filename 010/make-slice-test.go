package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 324)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 325)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	// new array is created after dumping old one
	x = append(x, 326)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
