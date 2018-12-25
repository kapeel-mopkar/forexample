package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2) // buffered channel- allows 1 value
	c <- 42 // doesnot work usually but works with buffered channels
	
	c <- 43
	fmt.Println(<-c) 
	fmt.Println(<-c) 
}
