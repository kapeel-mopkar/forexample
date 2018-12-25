package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//send
	go func(c chan<- int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}(c)

	//receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}
