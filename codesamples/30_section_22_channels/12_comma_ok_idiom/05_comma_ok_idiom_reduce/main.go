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
	close(q)

}
func received(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("even ", v)
		case v := <-o:
			fmt.Println("even ", v)
		case _, _ = <-q:
			return
		}
	}

}
