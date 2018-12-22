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

	fmt.Printf("%v %v, loves following flavors of ice-cream :\n", p1.fname, p1.lname)
	for _, val := range p1.icecreamFlavors {
		fmt.Printf("\t%v\n", val)
	}
	fmt.Printf("\n%v %v, loves following flavors of ice-cream :\n", p2.fname, p2.lname)
	for _, val := range p2.icecreamFlavors {
		fmt.Printf("\t%v\n", val)
	}

}