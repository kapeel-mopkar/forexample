package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x) //[4 5 6 7]
	x = append(x, 8, 9, 10)
	fmt.Println(x) // [4 5 6 7 8 9 10]

	y := []int{44, 55, 66, 77}
	x = append(x, y...) // Appending unfurled values from inside aggregate data structure
	fmt.Println(x) // [4 5 6 7 8 9 10 44 55 66 77]
	
	x = append(x[:2], x[4:]...) // removed 6 and 7 from slice
	fmt.Println(x) // [4 5 8 9 10 44 55 66 77]
}