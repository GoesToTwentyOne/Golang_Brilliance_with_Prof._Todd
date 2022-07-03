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
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("CPUs :", runtime.NumCPU())
	fmt.Println("Go Routine :", runtime.NumGoroutine())

}
