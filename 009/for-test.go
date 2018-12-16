package main

import (
	"fmt"
)

func main() {
	// incrementing i by 5--  for i := 0; i <= 100; i=(i+5) {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("Outer Loop : ", i)
		for j := 0; j < 3; j++ {
			fmt.Println("\tInner Loop : ", j)
		}
	}

	k := 1
	for k < 10 {
		fmt.Println(k)
		k++
	}

	l := 1
	for {
		if l > 9 {
			break
		}
		fmt.Println(l)
		l++
	}
}
