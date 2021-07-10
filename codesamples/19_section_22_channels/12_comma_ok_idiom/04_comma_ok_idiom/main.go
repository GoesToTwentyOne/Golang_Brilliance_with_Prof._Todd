package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 32
		close(ch)

	}()
	i, ok := <-ch
	fmt.Println(i, ok)
	i, ok = <-ch
	fmt.Println(i, ok)
}
