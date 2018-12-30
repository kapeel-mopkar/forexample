package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	const goroutes = 10

	for r := 0; r < goroutes; r++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
	
	for k := 0; k < 100; k++{
		fmt.Println(k, <-c)
	}
	close(c)
	fmt.Println("about to exit")
}
