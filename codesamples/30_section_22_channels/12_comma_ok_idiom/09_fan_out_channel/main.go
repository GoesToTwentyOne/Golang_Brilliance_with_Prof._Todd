package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go populate(ch1)
	go fanout(ch1, ch2)
	for v := range ch2 {
		fmt.Println(v)
	}
	fmt.Println("exit")
}
func populate(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i

	}
	close(ch1)

}
func fanout(ch1, ch2 chan int) {
	var wg sync.WaitGroup
	for v := range ch1 {

		wg.Add(1)
		go func(v1 int) {
			ch2 <- timeConsumingWork(v1)
			wg.Done()

		}(v)

	}
	wg.Wait()
	close(ch2)
}
func timeConsumingWork(v1 int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v1 + rand.Intn(1000)

}
