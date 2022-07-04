package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c1 := make(chan int, 1)
	c1 <- 22
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println(<-c1)
}
