package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumCPU())
	incrementor := 0
	gr := 100

	var wg sync.WaitGroup
	wg.Add(gr)
	var mu sync.Mutex
	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			v++
			incrementor = v
			mu.Unlock()
			wg.Done()

		}()
		fmt.Println("Goroutines ", runtime.NumCPU())
	}
	wg.Wait()
	fmt.Println(incrementor)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumCPU())

}
