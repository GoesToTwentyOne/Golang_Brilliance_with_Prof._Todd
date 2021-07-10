package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
	//buffered channel
	ch1 := make(chan int, 1)

	ch1 <- 42

	fmt.Println(<-ch1)

	//buffered channel
	ch2 := make(chan int, 1)
	go func() {
		ch2 <- 42
	}()
	fmt.Println(<-ch2)

}
