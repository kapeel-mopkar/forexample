package main

import (
	"fmt"
)

func main(){
	c := make(<- chan int, 1)

	//c <- 42 // invalid operation: c <- 42 (send to receive-only type <-chan int)
	
	//fmt.Println(<-c)
	fmt.Println("---------------")
	fmt.Printf("%T\n", c)

	c1 := make(chan <- int, 1)

	//c1 <- 42 
	
	//fmt.Println(<-c1) //invalid operation: <-c (receive from send-only type chan<- int)
	fmt.Println("---------------")
	fmt.Printf("%T\n", c1)


}