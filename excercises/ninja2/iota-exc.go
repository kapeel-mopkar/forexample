package main

import (
	"fmt"
)

const (
	a = (iota + 2018)
	b = (iota + 2018)
	c = (iota + 2018)
	d = (iota + 2018)
)

func main() {
	fmt.Println(a, b, c, d)
}