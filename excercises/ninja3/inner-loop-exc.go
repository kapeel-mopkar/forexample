package main

import (
	"fmt"
)

//Print every rune code point of the uppercase alphabet three times. Your output should look like this:
func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
