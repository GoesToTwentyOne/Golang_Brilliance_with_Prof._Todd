package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 21
	}()
	fmt.Println(<-ch)
}
