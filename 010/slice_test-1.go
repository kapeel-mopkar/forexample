package main

import (
	"fmt"
)

//slice allows values of same type togather
func main() {
	x:= []int{4, 5, 6, 7}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)
	fmt.Println(x[0]) // accessing value by index position
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	
	//range usage
	for i, v := range x {
		fmt.Println(i, v)
	}
}
