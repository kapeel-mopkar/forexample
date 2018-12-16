package main

import (
	"fmt"
)

func main() {
	switch {
	case (2 == 3):
		fmt.Println("Should not print")
	case (3 == 3):
		fmt.Println("Should print")

	}
}
