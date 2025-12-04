package main

import "fmt"

func main() {
	n := 5
	for i:=range n {
		for range i+1 {
			fmt.Print("*")
		}
		fmt.Println()

	}

}
