package main

import "fmt"

func main() {
	ch := make(chan int)
	go gosend(ch)
	goreceived(ch)
	fmt.Println("exit")
}
func gosend(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func goreceived(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
