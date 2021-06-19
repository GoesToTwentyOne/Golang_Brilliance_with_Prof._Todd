package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	counter := 0
	const rc = 100
	var wg sync.WaitGroup
	wg.Add(rc)
	var mu sync.Mutex

	for i := 0; i < rc; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()

		}()
		fmt.Println("Goroutine :", runtime.NumGoroutine())

	}
	wg.Wait()

	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Goroutine :", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
