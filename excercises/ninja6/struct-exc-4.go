package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func person speak(p person){
	fmt.Println(p.first, p.last)
}

func main(){
	p1 := person{
		first: "James",
		last: "Bond",
		age: 35,
	}

	p1.speak()
}