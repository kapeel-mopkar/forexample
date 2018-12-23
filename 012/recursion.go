package main

import (
	"fmt"
)

func main() {
	n := factorial(6)
	fmt.Println(n)
	
	t := loopFactorial(6)
	fmt.Println(t)
}

//recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//without recursion
func loopFactorial(n int) int {
	t:=1
	for i:=n; i>0; i-- {
		t*=i
	}
	return t
}
