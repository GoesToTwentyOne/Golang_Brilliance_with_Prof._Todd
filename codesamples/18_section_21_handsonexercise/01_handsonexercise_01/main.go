package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from goroutine 1")
		wg.Done()

	}()
	go func() {
		fmt.Println("Hello from goroutine 2")
		wg.Done()

	}()
	wg.Wait()
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Goroutines :", runtime.NumGoroutine())

}
