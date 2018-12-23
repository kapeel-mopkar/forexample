package main

import (
	"fmt"
)

func main(){
	f1:= func(xi ...int) int{
		total:=0
		for _,val:= range xi {
			total += val
		}
		return total
	}

	g1 := f1

	fmt.Println(g1(1,2,3,4,5,6,7,8,9))
}
