package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 32
	}()
	i, ok := <-ch
	fmt.Println(i, ok)

}
