package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 231
	ch <- 243
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
