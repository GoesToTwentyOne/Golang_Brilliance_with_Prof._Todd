package main

import "fmt"

func main() {
	ch := make(chan int)
	go send(ch)
	receive(ch)

	fmt.Println("all exit")
}
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

}
func receive(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)

	}

}
