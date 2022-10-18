package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup
	grtns := 100
	wg.Add( grtns )
	var count int

	var mu sync.Mutex
	
	for i := 0; i < grtns; i++ {
		go func(){
			mu.Lock()
			a := count
			runtime.Gosched()
			a++
			count = a
			fmt.Println( count)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println( count)
}