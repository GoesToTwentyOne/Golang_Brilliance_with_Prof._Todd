package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 21
	fmt.Println(<-ch)
}
