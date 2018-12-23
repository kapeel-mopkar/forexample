package main

import (
	"fmt"
)

func main(){
	fmt.Println(func(xi ...int) int{
		total:=0
		for _,val:= range xi {
			total += val
		}
		return total
	}(1,2,3,4,5,6,7,8,9))

}
