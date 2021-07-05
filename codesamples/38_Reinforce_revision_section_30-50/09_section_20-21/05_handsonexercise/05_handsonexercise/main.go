package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumCPU())
	var incrementor int64
	gr := 100

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Println("incermentor ato :", atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
		fmt.Println("Goroutines ", runtime.NumCPU())
	}
	wg.Wait()
	fmt.Println(incrementor)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumCPU())

}
