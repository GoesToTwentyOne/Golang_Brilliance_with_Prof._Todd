package main

import "fmt"

func main() {
	ch := make(chan<- int, 2)
	ch <- 2
	ch <- 23
	//doesn't run
	fmt.Println(<-ch)
}
