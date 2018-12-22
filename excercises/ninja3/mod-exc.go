package main

import (
	"fmt"
)
//Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
func main() {
	for i:=10; i<=100; i++{
		if(i%4 != 0)
		fmt.Println("Number : ",i, ", Remainder modulus of 4 : ", i%4)
	}
}
