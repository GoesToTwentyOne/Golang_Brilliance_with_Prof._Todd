package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())
	inc := 0
	gr := 50
	var wg sync.WaitGroup

	wg.Add(gr)
	var mu sync.Mutex
	for i := 0; i < gr; i++ {

		go func() {
			mu.Lock()

			v := inc
			v++
			inc = v

			fmt.Println("\t\tincrement2", inc)
			mu.Unlock()
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())
	fmt.Println("\t\t\t\tincrement3", inc)

}
