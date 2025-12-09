package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	// role := "worker"

	// switch role {
	// case "admin":
	// 	fmt.Println("all access")
	// 	fallthrough
	// case "emp":
	// 	fmt.Println("read edit")
	// 	fallthrough
	// case "wokrker":
	// 	fmt.Println("read")
	// default:
	// 	fmt.Println("no accesss")
	// }

	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(100))
	d2 := time.Duration(rand.Int63n(100))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 11
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 22

	}()

	select {
	case v1 := <-ch1:
		fmt.Println("got from ch1", v1)
	case v2 := <-ch2:
		fmt.Println("got from ch2", v2)
	}
}
