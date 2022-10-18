package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 43

	fmt.Println(<-c)

}