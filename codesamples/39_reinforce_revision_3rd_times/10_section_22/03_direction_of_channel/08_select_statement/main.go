package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	received(even, odd, quit)

}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0

}
func received(e, o, q <-chan int) {
	for i := 0; i < 100; i++ {
		select {
		case <-e:
			fmt.Println("Even :", i)
		case <-o:
			fmt.Println("Odd :", i)
		case <-q:
			fmt.Println("quit :", i)
		default:
			fmt.Println("From Default")
		}

	}
}
