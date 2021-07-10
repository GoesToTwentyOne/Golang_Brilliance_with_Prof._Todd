package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var inc int64
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("CPUs : ", runtime.NumGoroutine())

	rc := 50

	wg.Add(rc)
	for i := 0; i < rc; i++ {
		go func() {

			fmt.Println("add ", atomic.AddInt64(&inc, 1))
			fmt.Println("increment ", atomic.LoadInt64(&inc))

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("CPUs : ", runtime.NumGoroutine())
	fmt.Println("final increment ", inc)
}
