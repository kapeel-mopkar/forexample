package main

import ("fmt")
var y string
var z int
func main(){
	fmt.Println("Value of Y ;'", y, "'")
	fmt.Printf("%T", y)

	y= "Bond, James"

	fmt.Println("Value of Y ;'", y, "'", z)
	fmt.Printf("%T", y)
}