package main

import (
	"fmt"
)

func main() {
	jb := []string{"james", "bond", "chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "moneypenny", "strawberry", "hazelnut"}
	fmt.Println(mp)
	
	//slice of slice / multidimensional slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
