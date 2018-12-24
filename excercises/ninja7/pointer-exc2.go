package main

import (
	"fmt"
)

type person struct{
	first string
	last string
	age int
}

func main(){
	p1 := person{
		first: "James",
		last: "Bond",
		age: 35,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person){
	fmt.Println(*p)
	//p.age=38 // This works as well
	(*p).age=36
	fmt.Println(*p)
}