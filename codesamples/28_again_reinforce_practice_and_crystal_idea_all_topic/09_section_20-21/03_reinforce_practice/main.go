package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("cpus ", runtime.NumCPU())
	fmt.Println("goroutines ", runtime.NumGoroutine())
	counter := 0
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("cpus ", runtime.NumCPU())
	fmt.Println("goroutines ", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
