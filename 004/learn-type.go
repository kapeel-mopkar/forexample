package main

import ("fmt")

var y=42
var z= "Shaken, but not stirred!"

// back quote not single quote, for string literals
var a= `James said, 
"Shaken, but not stirred
!!!"`

func main(){
	fmt.Println(y);
	fmt.Printf("%T\n", y)

	fmt.Println(z);
	fmt.Printf("%T\n", z)

	fmt.Println(a);
	fmt.Printf("%T\n", a)

	/*
	// Static programming language, not a dynamic prog language
	// Variables can hold a value of only certain type with which its instantiated
	z=43
	fmt.Println(z);
	fmt.Printf("%T\n", z)*/

}