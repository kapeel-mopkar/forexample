package main

import (
	"fmt"
)

func main() {
	//initialization statement
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	//fmt.Println(x) //-- cant print x as its not on scope // ERROR: prog.go:12:15: undefined: x
	fmt.Println("statement 1")
	fmt.Println("statement 2")

}
