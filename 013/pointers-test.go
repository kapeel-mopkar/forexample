package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	
	b := &a // var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // gives value stored at an address
	
	fmt.Println(*&a)
	
	*b = 43 // value stored at address of a is being changed
	fmt.Println(a)
	
}
