package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	d = iota // ease of programming, do not have to repeat!!!
	e
	f
)

func main() {
	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d)

	fmt.Println(e)

	fmt.Println(f)
}