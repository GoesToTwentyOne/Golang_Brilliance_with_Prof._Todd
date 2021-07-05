package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutine ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("execurint from 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("execurint from 2")
		wg.Done()
	}()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutine ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("All are done")

}
