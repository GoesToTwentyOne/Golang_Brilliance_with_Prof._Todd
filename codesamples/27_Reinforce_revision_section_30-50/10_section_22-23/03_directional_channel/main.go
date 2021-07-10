package main

import "fmt"

func main() {
	ch := make(chan int)
	//send
	go chsend(ch)
	// received
	chreceived(ch)
	fmt.Println("done")

}

func chsend(c chan<- int) {
	c <- 42
}
func chreceived(c <-chan int) {

	fmt.Println(<-c)
}
