package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			ch <- i
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Println(value)
	}
}
