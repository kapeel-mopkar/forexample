package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent {
		person: person{
			fname: "James",
			lname: "Bond",
			age: 32,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	
	fmt.Println(sa1.fname, sa1.lname, sa1.age, sa1.ltk)

}
