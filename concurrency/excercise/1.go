package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hi 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("hi 2")
		wg.Done()
	}()

	wg.Wait()
}