package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println((*p).name, (*p).age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Kapeel Mopkar",
		age:  35,
	}

	saySomething(&p1)
	//saySomething(p1) // Does not work
	p1.speak() // This works, it directly calls the speak method, which has receiver is a poiter to person, but person works
}
