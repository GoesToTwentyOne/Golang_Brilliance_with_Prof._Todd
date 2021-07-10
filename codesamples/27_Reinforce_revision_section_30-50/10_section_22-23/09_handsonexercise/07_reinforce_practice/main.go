package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	fanout := make(chan int)
	go send(ch)
	go fanoutev(ch, fanout)
	for v := range fanout {
		fmt.Println(v)
	}
}
func send(c chan<- int) {

	go func() {
		for i := 0; i < 10; i++ {
			c <- i

		}
		close(c)

	}()
}
func fanoutev(c <-chan int, fan chan<- int) chan<- int {
	var wg sync.WaitGroup
	for v := range c {

		wg.Add(v)
		go func(n int) {
			fmt.Println("Goroutine :", n)
			for i := 0; i < 10; i++ {
				fan <- random(n)
				wg.Done()

			}
			close(fan)

		}(v)

	}

	wg.Wait()
	return fan

}
func random(n int) int {
	return n + rand.Intn(1000)

}
