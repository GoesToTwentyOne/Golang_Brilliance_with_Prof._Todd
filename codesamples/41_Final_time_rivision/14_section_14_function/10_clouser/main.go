package main

import "fmt"

func main() {
	a := increment()()
	fmt.Println(a)
	b := increment()()
	fmt.Println(b)
	c := increment()()
	fmt.Println(c)

}
func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
