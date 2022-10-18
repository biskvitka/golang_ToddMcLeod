package main

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	grtns := 100
	wg.Add( grtns )
	var count int64

	for i := 0; i < grtns; i++ {
		go func(){
			runtime.Gosched()
			atomic.AddInt64( &count, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println( count)
}