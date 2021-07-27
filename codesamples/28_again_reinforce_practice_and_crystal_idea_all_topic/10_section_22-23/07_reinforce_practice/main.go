package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	//send
	go send(chan1)
	//received
	go fanout(chan1, chan2)
	for v := range chan2 {
		fmt.Println(v)
	}
	fmt.Println("all exit")
}
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)

}
func fanout(c <-chan int, fanout chan<- int) {
	var wg sync.WaitGroup
	for v := range c {
		wg.Add(1)
		func(v2 int) {
			fanout <- timeConsumingWork(v2)
			wg.Done()

		}(v)
	}
	wg.Wait()
	close(fanout)

}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
