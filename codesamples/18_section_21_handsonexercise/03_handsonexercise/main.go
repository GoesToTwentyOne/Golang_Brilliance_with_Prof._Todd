package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	inc := 0
	gr := 100
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {

			v := inc
			runtime.Gosched()
			v++
			inc = v
			fmt.Println("increment1", inc)
			wg.Done()

		}()
		fmt.Println("\tincrement2", inc)
	}
	wg.Wait()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("\t\tincrement3", inc)

}
