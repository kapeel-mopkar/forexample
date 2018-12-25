package main

import (
	"fmt"
	"sync"
)

func main(){
	loop := 2
	var wg sync.WaitGroup
	wg.Add(loop)
	for i:=0; i < loop; i++{
		go func() {
			fmt.Println("Testing Go Routines")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Completed Testing Go Routines")
}