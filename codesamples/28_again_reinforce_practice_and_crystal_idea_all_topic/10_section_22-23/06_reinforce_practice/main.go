package main

import (
	"fmt"
	"sync"
)

func main() {
	chEven := make(chan int)
	chOdd := make(chan int)
	fanin := make(chan int)
	//send
	go send(chEven, chOdd)

	//received
	go received(chEven, chOdd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("all exit")
}
func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)

}
func received(e, o <-chan int, fanint chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fanint <- v

		}
		wg.Done()

	}()

	go func() {
		for v := range o {
			fanint <- v

		}
		wg.Done()
	}()
	wg.Wait()
	close(fanint)

}
