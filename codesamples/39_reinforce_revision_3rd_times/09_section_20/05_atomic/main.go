package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var counter int64 = 0
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
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
