package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gr = 100
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count: ", counter)

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines END", runtime.NumGoroutine())

}
