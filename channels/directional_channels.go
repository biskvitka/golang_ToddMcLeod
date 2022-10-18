package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int, 2) // send only
	cr := make(<-chan int, 2) // receive only
	c := make(chan int, 2) // bidirectional channel
	c <- 43 // send
	c <- 443

	fmt.Println(<-c) // receive
	fmt.Println(<-c)
	fmt.Println("---------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general channel assignment/conversion is not possible
	// c = cr
	// c = cs

	// general to specific assignment/conversion is possible
	cs = c
	cr = c
}