package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	go send(even, odd, quit)
	received(even, odd, quit)

}
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)

}
func received(e, o <-chan int, q <-chan bool) {
	for i := 0; i < 100; i++ {
		select {
		case <-e:
			fmt.Println("Even :", i)
		case <-o:
			fmt.Println("Odd :", i)
		case i, ok := <-q:
			if !ok {
				fmt.Println("quit :", ok, i)
				return

			} else {
				fmt.Println("quit :", ok, i)
			}
		}

	}
}
