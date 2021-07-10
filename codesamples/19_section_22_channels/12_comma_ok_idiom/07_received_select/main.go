package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		for i := 0; i < 5000; i++ {
			a <- i

		}

	}()

	go func() {
		for i := 0; i < 5000; i++ {
			b <- i

		}

	}()

	received(a, b)

}
func received(a, b chan int) {
	for i := 0; i < 10000; i++ {
		select {
		case v := <-a:
			fmt.Println("a ", v)

		case v := <-b:
			fmt.Println("b ", v)
		}

	}

}
