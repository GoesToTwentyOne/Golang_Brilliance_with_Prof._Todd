package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOOS :\t\t", runtime.GOOS)
	fmt.Println("GOARCH :\t\t", runtime.GOARCH)
	fmt.Println("CPU :\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINE :\t\t", runtime.NumGoroutine())
}
