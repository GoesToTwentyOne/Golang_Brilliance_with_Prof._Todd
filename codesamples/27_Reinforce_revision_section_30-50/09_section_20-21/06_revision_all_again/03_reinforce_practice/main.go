package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	incrementor := 0
	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			va := incrementor
			va++
			runtime.Gosched()
			incrementor = va
			wg.Done()
		}()
		fmt.Println("Goroutines func ", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Total incrementor ", incrementor)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

}
