package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := bar()()
	fmt.Println(s2)
}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
