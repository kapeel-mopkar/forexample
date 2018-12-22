package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x)
	x = append(x, 8, 9, 10)
	fmt.Println(x)

	y := []int{44, 55, 66, 77}
	x = append(x, y...) // Appending unfurled values from inside aggregate data structure
	fmt.Println(x)
}
