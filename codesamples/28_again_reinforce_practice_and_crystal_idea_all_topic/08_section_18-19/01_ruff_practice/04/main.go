package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("before Goroutines", runtime.NumGoroutine())
	fmt.Println("before Num of cpus", runtime.NumCPU())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("after Goroutines", runtime.NumGoroutine())
	fmt.Println("after Num of cpus", runtime.NumCPU())
	wg.Wait()

}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("I am from foo      index ", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("I am from bar      index ", i)
	}

}
