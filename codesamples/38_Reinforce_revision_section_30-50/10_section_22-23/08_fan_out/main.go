package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go send(c1)
	go fanout(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}
func send(c chan int) {

	for i := 0; i < 10; i++ {
		c <- i

	}

	close(c)

}
func fanout(c <-chan int, cf chan<- int) {
	var wg sync.WaitGroup

	for v := range c {
		wg.Add(1)
		go func(n int) {
			cf <- randomnub(n)
			wg.Done()

		}(v)

	}

	wg.Wait()
	close(cf)

}
func randomnub(n int) int {

	return n + rand.Intn(1000)
}
