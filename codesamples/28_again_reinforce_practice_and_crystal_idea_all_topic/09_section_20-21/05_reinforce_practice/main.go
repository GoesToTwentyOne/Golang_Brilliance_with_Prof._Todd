package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var increment int64
	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			fmt.Println(atomic.LoadInt64(&increment))
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println("count:", increment)

}
