package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	const gr = 100
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter \t", atomic.LoadInt64(&counter))
			wg.Done()

		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("main counter ", counter)

}
