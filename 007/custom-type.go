package main

import ("fmt")

var a int
type Pizza int
var b Pizza
func main(){
	a = 42
	b=43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Conversion not casting
	a=int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}