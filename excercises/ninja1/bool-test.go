package main

import ("fmt")

var x bool

func main(){
	fmt.Println(x)
	x=true
	fmt.Println(x)
	a:=2
	b:=3
	fmt.Println(a==b)
	a=b
	fmt.Println(a==b, a != b)
}