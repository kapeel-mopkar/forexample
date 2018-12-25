package main

import (
	"fmt"
)

func main(){
	c := make(chan int)

	//send
	go send(c)

	//receive
	receive(c)

	fmt.Println("About to exit")

}

//send
func send(c chan<- int){
	c <- 42
}

//receive
func receive(c <-chan int){
	fmt.Println("Received :", <-c)
}