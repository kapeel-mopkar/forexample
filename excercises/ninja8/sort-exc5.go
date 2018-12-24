package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

type ByLast []user

type BySayings []string

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByLast) Len() int {
	return len(a)
}

func (a ByLast) Less(i, j int) bool {
	return a[i].Last < a[j].Last
}

func (a ByLast) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a BySayings) Len() int {
	return len(a)
}

func (a BySayings) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a BySayings) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	fmt.Println(">>>>>>>> By INSERTION ORDER >>>>>>>")

	for i, val := range users {
		fmt.Printf("\n%d) First Name='%v', Last Name='%v', Aged='%v' says >\n", i, val.First, val.Last, val.Age)
		for j, say := range val.Sayings {
			fmt.Printf("\t%d) '%v'\n", j, say)
		}
	}

	sort.Sort(ByAge(users))
	//fmt.Println(users)
	fmt.Println("\n>>>>>>>> Sorted By AGE >>>>>>>")
	for i, val := range users {
		fmt.Printf("\n%d) First Name='%v', Last Name='%v', Aged='%v' says >\n", i, val.First, val.Last, val.Age)
		for j, say := range val.Sayings {
			fmt.Printf("\t%d) '%v'\n", j, say)
		}
	}

	sort.Sort(ByLast(users))
	//fmt.Println(users)
	fmt.Println("\n>>>>>>>> Sorted By Last Name >>>>>>>")
	for i, val := range users {
		fmt.Printf("\n%d) First Name='%v', Last Name='%v', Aged='%v' says >\n", i, val.First, val.Last, val.Age)
		for j, say := range val.Sayings {
			fmt.Printf("\t%d) '%v'\n", j, say)
		}
	}

	fmt.Println("\n>>>>>>>> Sorted By Last Name, Sorted Sayings >>>>>>>")
	//fmt.Println(users)
	for i, val := range users {
		fmt.Printf("\n%d) First Name='%v', Last Name='%v', Aged='%v' says >\n", i, val.First, val.Last, val.Age)
		sort.Sort(BySayings(val.Sayings))
		for j, say := range val.Sayings {
			fmt.Printf("\t%d) '%v'\n", j, say)
		}
	}

}
