package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("GOARCH", runtime.GOARCH)
	fmt.Println("NUM CUPs", runtime.NumCPU())
	fmt.Println("NUM Goroutines", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("NUM CUPs", runtime.NumCPU())
	fmt.Println("NUM Goroutines", runtime.NumGoroutine())
	wg.Wait()
}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foor  ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar  ", i)
	}
}
