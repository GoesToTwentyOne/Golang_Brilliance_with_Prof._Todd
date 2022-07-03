package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Go Routines :", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Func 1")
		wg.Done()

	}()
	go func() {
		fmt.Println("Func 2")
		wg.Done()

	}()

	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Go Routines :", runtime.NumGoroutine())
	wg.Wait()

}
