package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
	Area() float64
}

func getArea(s Shape) float64 {
	return s.Area()
}

func run_interface1() {

	rad := Rectangle{3, 4}
	cir := Circle{3}
	fmt.Println(getArea(rad))
	fmt.Println(getArea(cir))

}
