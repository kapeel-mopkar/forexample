package main

import ("fmt")

type keychain int
var x keychain
var y int
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y=int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}