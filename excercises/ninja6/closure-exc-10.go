package main

import (
	"fmt"
)

func main(){
	a:= foo()
	b:= foo()

	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())

	a= foo()
	b= foo()

	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
}

func foo() func() int{
	x := 0
	return func() int {
		x++
		return x
	}
}