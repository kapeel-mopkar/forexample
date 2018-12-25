package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	
	var incrementer int64
	gs := 100
	wg.Add(gs)
	fmt.Println("Before Routine:", incrementer)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			runtime.Gosched()
			fmt.Println("\t In Routine", atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("After Routine:", incrementer)
}