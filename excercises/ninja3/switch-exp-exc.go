package main

import (
	"fmt"
)

//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
func main() {
	favSport := "Cricket"
	switch favSport {
	case "Cricket":
		fmt.Println("I love Cricket")
	case "Football":
		fmt.Println("I love Football")
	default:
		fmt.Println("I hate sports")
	}
}
