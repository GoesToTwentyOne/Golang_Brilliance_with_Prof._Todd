package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 121
	x[9] = 43
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 455444)
	x = append(x, 455445)
	x = append(x, 455446)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
