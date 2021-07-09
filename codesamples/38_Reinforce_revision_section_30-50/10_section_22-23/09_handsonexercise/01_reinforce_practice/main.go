package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 21
	}()
	fmt.Println(<-ch)

	ch1 := make(chan int, 1)
	ch1 <- 22
	fmt.Println(<-ch1)
}
