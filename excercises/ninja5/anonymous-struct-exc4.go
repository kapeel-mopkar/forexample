package main

import ("fmt")

type vehicle struct {
	doors int
	color string
}

type truck struct{
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main(){
	//Anonymous Strucs
	t := struct{
		vehicle
		fourWheel bool
		}{
		vehicle: vehicle{
			doors: 5,
			color: white,
		},
		fourWheel: true
	}

	s := struct{
		vehicle
		luxury bool
		}{
		vehicle: vehicle{
			doors: 4,
			color: black,
		},
		luxury: true
	}

	fmt.Println(t)
	fmt.Println(s)
}