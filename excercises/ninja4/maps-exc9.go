package main

import (
	"fmt"
)

func main() {
	m1 := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m1["mopkar_kapeel"] = []string{`Programming/Coding`, `Sci-fi shows`, `Video games`}
	for k, v := range m1 {
		fmt.Println("Key : ", k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}

}
