package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())

}
