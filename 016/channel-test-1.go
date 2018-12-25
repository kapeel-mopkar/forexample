package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // buffered channel- allows 1 value
	c <- 42 // doesnot work usually but works with buffered channels
	
	// c <- 43 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c) 
}
