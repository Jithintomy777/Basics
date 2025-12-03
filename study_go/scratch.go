package main

import "fmt"

func runScratch() {

	q := []int{12, 22, 32}
	sum := 0
	for _, i := range q {
		sum = sum + i
		fmt.Println(sum)
	}

}
