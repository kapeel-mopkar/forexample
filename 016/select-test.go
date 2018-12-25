package main

import (
	"fmt"
)

func main(){
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
	fmt.Println("About to exit")
}

func receive(e, o, q <-chan int){
	for {
		select {
			case v:= <- e:
				fmt.Println("From Eve channel:", v)
			case v:= <- o:
				fmt.Println("From Odd channel:", v)
			case v:= <- q:
				fmt.Println("From Quit channel:", v)
				//close(q)
				return
		}
	}
}

func send(e, o, q chan<- int){
	for i:=0; i<50; i++{
		if i%2 == 0 {
			e<-i
		}
	}
	for i:=0; i<50; i++{
		if i%2 != 0 {
			o<-i
		}
	}
	//close(e)
	//close(o)
	q <- 0
}