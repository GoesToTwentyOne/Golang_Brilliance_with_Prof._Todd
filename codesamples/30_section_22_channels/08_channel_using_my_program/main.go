package main

import "fmt"

func main() {
	ch := make(chan int)
	//send
	go goes(ch)
	//received
	goe(ch)
	fmt.Println("exit")
}

//send
func goes(ch chan<- int) {
	ch <- 45
}

//received
func goe(ch <-chan int) {
	fmt.Println(<-ch)
}
