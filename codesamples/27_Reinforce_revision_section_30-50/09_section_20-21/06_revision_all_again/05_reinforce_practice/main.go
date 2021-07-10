package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var incrementor int64 = 0
	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {

			atomic.AddInt64(&incrementor, 1)

			fmt.Println("incrementor inside ano", atomic.LoadInt64(&incrementor))

			wg.Done()
		}()
		fmt.Println("Goroutines func ", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Total incrementor ", incrementor)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

}
