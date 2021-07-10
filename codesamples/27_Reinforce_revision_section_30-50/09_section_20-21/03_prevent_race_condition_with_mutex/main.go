package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines ", runtime.NumGoroutine())
	conter := 0
	const goroutinesforProgram = 100
	var wg sync.WaitGroup
	wg.Add(goroutinesforProgram)
	var mu sync.Mutex
	for i := 0; i < goroutinesforProgram; i++ {
		go func() {

			mu.Lock()
			v := conter
			runtime.Gosched()
			v++
			conter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("count: ", conter)

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines End", runtime.NumGoroutine())

}
