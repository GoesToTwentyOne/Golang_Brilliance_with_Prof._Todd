package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case v, ok := <-q:
			if ok {
				fmt.Println(v)
				return
			} else {
				fmt.Println(v)
				return

			}
		}

	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
		close(c)
	}()

	return c
}
