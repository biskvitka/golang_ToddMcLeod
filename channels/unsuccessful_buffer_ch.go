package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 43
	c <- 443

	fmt.Println(<-c)
}