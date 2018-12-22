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
	//Normal struct, which is defined above
	p1 := person{
		fname: "James",
		lname: "Bond",
		age:   32,
	}
	
	//Anonymous struct 
	p2 := struct {
		fname string
		lname string
		age   int
	}{
		fname: "Miss",
		lname: "Moneypenny",
		age:   30,
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
