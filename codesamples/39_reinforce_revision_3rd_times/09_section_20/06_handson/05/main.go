package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())

}
