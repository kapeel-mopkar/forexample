package main

import ("fmt")

var x int
var y float64 
var z int8 = -128 // range -128 to 127 8 bit signed int
var p byte
func main(){
	x = 42
	y = 42.336
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	p=34
	fmt.Println(p)
	fmt.Printf("%T\n", p)

}