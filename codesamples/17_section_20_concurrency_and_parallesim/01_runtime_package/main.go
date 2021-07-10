package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go architecher ", runtime.GOARCH)
	fmt.Println("Go operating system ", runtime.GOOS)
	fmt.Println("Number of cpus ", runtime.NumCPU())
	fmt.Println("Number of Goroutines ", runtime.NumGoroutine())
}
