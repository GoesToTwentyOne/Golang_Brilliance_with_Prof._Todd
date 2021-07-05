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
	for i := 0; i < gr; i++ {
		go func() {
			v := incrementor
			v++
			runtime.Gosched()
			incrementor = v
			wg.Done()

		}()
		fmt.Println("Goroutines ", runtime.NumCPU())
	}
	wg.Wait()
	fmt.Println(incrementor)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumCPU())

}
