package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	fmt.Println("GOOS :", runtime.GOOS)
	fmt.Println("GOARH :", runtime.GOARCH)
	wg.Add(2)
	go sum()
	go multiplication()

	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Done")

}
func sum() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum : ", sum)
	wg.Done()
}
func multiplication() {
	mul := 1
	for i := 1; i < 10; i++ {
		mul *= i
	}
	fmt.Println("mul : ", mul)
	wg.Done()
}
