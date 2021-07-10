package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 32
		ch <- 33

	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
