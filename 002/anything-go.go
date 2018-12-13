package main

import "fmt"

func main() {
	n, err := fmt.Println("Main Test")
	foo()
	fmt.Println("Foo End")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
	fmt.Println("Main End", n, err)
}

func foo() {
	fmt.Println("Foo Test")
}

func bar() {
	fmt.Println("Bar Test", 101, true)
}
