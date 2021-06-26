package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4))

}
func sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
