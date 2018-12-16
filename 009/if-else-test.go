package main

import (
	"fmt"
)

func main() {
	x := 41

	if x == 40 {
		fmt.Println("Value is 40")
	} else if x == 41 {
		fmt.Println("Value is 41")
	} else {
		fmt.Println("Value is not 40 or 41")
	}

}