package main

import ("fmt")

type keychain int
var x keychain

func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x:=42
	fmt.Println(x)
}