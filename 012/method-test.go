package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{fname: "James",
			lname: "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	sa2 := secretAgent{
		person: person{fname: "Miss",
			lname: "Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa2)
	sa2.speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.ltk, s.fname, s.lname)
}
