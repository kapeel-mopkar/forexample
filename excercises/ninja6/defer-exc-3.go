package main

import (
	"fmt"
)

func main(){
	defer foo()
	bar()
	fmt.Println("Main completed")
}

func foo(){
	defer func(){
		fmt.Println("Anonymous completed")
	}
	fmt.Println("Foo completed")
}

func bar(){
	fmt.Println("Bar completed")
}