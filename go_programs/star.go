package main

import "fmt"

func star() {
	n := 5
	for i := range n {
		for range i + 1 {
			fmt.Print("*")
		}
		fmt.Println()

	}

}
