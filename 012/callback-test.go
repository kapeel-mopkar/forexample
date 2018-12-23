package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("Total:", s)

	s2 := even(sum, ii...)
	fmt.Println("Even total: ", s2)

	s3 := odd(sum, ii...)
	fmt.Println("Odd total: ", s3)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var xe []int

	for _, v := range vi {
		if v%2 == 0 {
			xe = append(xe, v)
		}
	}

	return f(xe...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var xo []int

	for _, v := range vi {
		if v%2 != 0 {
			xo = append(xo, v)
		}
	}

	return f(xo...)
}
