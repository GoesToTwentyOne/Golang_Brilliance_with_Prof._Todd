package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	go receive(c)

}
func send(c chan<- int) {
	c <- 21
}
func receive(c <-chan int) {
	fmt.Println(<-c)
}
