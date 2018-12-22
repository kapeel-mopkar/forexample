package main

import (
	"fmt"
)

type person struct {
	fname           string
	lname           string
	icecreamFlavors []string
}

func main() {
	p1 := person{
		fname: "James",
		lname: "Bond",
		icecreamFlavors: []string{
			"Dark Chocolate",
			"Poison Ivy",
			"Chili Blast",
		},
	}

	p2 := person{
		fname: "Miss",
		lname: "Moneypenny",
		icecreamFlavors: []string{
			"Strawberry",
			"Honey Lemon",
			"Choco Vanilla",
		},
	}

	m := map[string]person{p1.lname: p1,
		p2.lname: p2,
	}

	for k, p := range m {
		fmt.Println(k, " likes :")
		for i, val := range p.icecreamFlavors {
			fmt.Printf("\t%v) %v\n", i+1, val)
		}
	}

}