package main
import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

//receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
func send(even, odd <-chan int, fanin chan<- int) {
	var wg sync.Waitgroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}

	wg.Wait()
	close(fanin)
}