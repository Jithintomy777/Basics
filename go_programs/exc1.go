package main

import "fmt"

var x int = 23

const y = 3

func exc1() {
	z := 12
	fmt.Printf("x value %v, type %T\n", x, x)
	fmt.Printf("y value %v, type %T\n", y, y)
	fmt.Printf("z value %v, type %T\n", z, z)
}
