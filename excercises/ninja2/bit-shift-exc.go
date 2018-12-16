package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("%d\t%b\t%#x", a, a, a)
	b := a << 1
	fmt.Println("%d\t%b\t%#x", b, b, b)
}
