package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int{
		"James": 32,
		"Penny": 27,
		"Kaps":  35,
	}

	v1 := m1["James"]
	fmt.Println(v1) // 32

	v2 := m1["james"]
	fmt.Println(v2) // 0

	v3, ok := m1["james"]
	fmt.Println(v3, ok) // 0 false

	if v4, ok := m1["James"]; ok {
		fmt.Println("James age is ", v4) // James age is  32
	}

	m1["Todd"] = 34

	for k, v := range m1 {
		fmt.Println("\t", k, v)
		/*
					Penny 27
			 		Kaps 35
			 		Todd 34
			 		James 32
		*/

	}

	xi := []int{4, 5, 6, 7, 8, 91}

	for i, v := range xi {
		fmt.Println("\t", i, v)
		/*
			0 4
	 		1 5
	 		2 6
	 		3 7
	 		4 8
	 		5 91
		*/
	}
}