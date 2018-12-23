package main

import (
	"fmt"
)

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x, y, z)
}

func foo() int {
	a := 42
	return a
}

func bar() (int, string) {
	i := 42
	s := "Bar String"

	return i, s
}
