package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 43
	}()
	fmt.Println(<-ch)

}
