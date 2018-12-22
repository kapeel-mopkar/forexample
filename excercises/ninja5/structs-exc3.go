package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 12345,
			"Dr Evil":    131313,
			"Prof Bravo": 999999,
		},
		favDrinks: []string{"martini", "water"},
	}

	fmt.Println(s)
	fmt.Println(s.first)
	fmt.Println(s.friends)
	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	fmt.Println(s.favDrinks)
	for i, v := range s.favDrinks {
		fmt.Println(i, v)
	}
}
