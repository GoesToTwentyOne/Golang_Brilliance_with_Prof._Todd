package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

	}()
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(<-c)
		}

	}()
}
