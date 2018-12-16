package main

import (
	"fmt"
)

func main() {
	x := "f"
	if x == "Bond, James" {
		fmt.Println("Hello, James")
	} else if x == "SOmething else" {
		fmt.Println("Hello, SOmething")
	} else {
		fmt.Println("Hello, Others")
	}
}
