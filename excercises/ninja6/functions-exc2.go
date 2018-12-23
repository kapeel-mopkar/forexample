package main

import (
	"fmt"
)

func main() {
	xi:=[]int{1,2,4,5,6,7,8,9,}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(numbers ...int) int{
	total :=0
	for _,val := range numbers{
		total += val
	}
	return total
}

func bar(xi []int) int {
	total :=0
	for _,val := range xi{
		total += val
	}
	return total
}