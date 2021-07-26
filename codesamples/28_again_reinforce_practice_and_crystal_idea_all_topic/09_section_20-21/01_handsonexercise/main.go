package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("arch", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("goroutines ", runtime.NumGoroutine())
	wg.Add(2)

	go foo()
	go bar()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("goroutines ", runtime.NumGoroutine())
	wg.Wait()

}
func foo() {
	for i := 0; i < 100; i++ {
		func() {
			fmt.Println("from foo", i)
		}()
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 100; i++ {
		func() {
			fmt.Println("from bar", i)
		}()
	}
	wg.Done()
}
