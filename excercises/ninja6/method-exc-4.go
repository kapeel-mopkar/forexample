package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Println(p.first, p.last, "is aged", p.age)
}

func main(){
	p1 := person{
		first: "James",
		last: "Bond",
		age: 35,
	}
	
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 30,
	}

	p1.speak()
	p2.speak()
}