package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

func main() {
	p1 := person{
		fname: "James",
		lname: "Bond",
		age:   32,
	}

	p2 := person{
		fname: "Miss",
		lname: "Moneypenny",
		age:   30,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println(p2.fname, p2.lname, p2.age)
}
