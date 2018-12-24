package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Fname string
	Lname string
	Age   int
	Addrs []address
}

type address struct {
	City        string
	State       string
	Zip         int
	Residential bool
}

func main() {
	addresses := []address{
		{
			City:        "Shiroda",
			State:       "Goa",
			Zip:         403706,
			Residential: true,
		}, {
			City:        "Margao",
			State:       "Goa",
			Zip:         403722,
			Residential: false,
		},
	}

	p1 := person{
		Fname: "Kapeel",
		Lname: "Mopkar",
		Age:   35,
		Addrs: addresses,
	}

	j, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Error !!!! - ", err)
	}
	
	fmt.Println(string(j))

	os.Stdout.Write(j)
	
	
}
