package main

import "fmt"

func main() {
	ch := make(chan<- int)
	go func() {
		ch <- 231
	}()
	// fmt.Println(<-ch)
	//error as send means send not received
	fmt.Printf("%T \n", ch)
}
