package main

import (
	"fmt"
)

func main() {

	f1 := foo()

	x:=f1()
	
	fmt.Println(x)

}

func foo() func() int {
	return func() int {
		return 42
	}
}
