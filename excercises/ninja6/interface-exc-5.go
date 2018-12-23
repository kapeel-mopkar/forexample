package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	radius float64
}

type CIRCLE struct {
	length  float64
}

type SHAPE interface {
	area() float64
}

func (s SQUARE) area() float64 {
	return (math.Pi * (s.radius * s.radius))
}

func (c CIRCLE) area() float64 {
	return (c.length * c.length)
}

func main() {
	s1 := SQUARE{
		radius: 12.345,
	}
	fmt.Println("SQUARE Area (in Sq. Mt.) :", s1.area())

	c1 := CIRCLE{
		length:  15.0,
	}
	fmt.Println("CIRCLE Area (in Sq. Mt.) :", c1.area())

	info(s1)
	info(c1)
}

func info(shape SHAPE) {
	fmt.Println("INFO Area (in Sq. Mt.) :", shape.area())
}
