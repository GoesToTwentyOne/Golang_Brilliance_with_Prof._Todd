package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()

		}()
		fmt.Println("Goroutine : ", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

}
