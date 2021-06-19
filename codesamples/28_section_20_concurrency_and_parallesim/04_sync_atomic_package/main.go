package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("goroutine ", runtime.NumGoroutine())
	var counter int64
	const rc = 100
	var wg sync.WaitGroup
	wg.Add(rc)

	for i := 0; i < rc; i++ {
		go func() {
			fmt.Println("count ", atomic.AddInt64(&counter, 1))
			runtime.Gosched()
			fmt.Println("count Read ", atomic.LoadInt64(&counter))
			wg.Done()

		}()
		fmt.Println("goroutine ", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("goroutine ", runtime.NumGoroutine())
	// fmt.Println("count2 ", counter)

}
