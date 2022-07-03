package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	received(c)
}
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

}
func received(c <-chan int) {

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

}
