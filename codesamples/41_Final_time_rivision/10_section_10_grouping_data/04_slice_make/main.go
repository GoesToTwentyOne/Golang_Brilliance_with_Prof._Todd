package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)
	fmt.Println("capacity :", cap(x))
	fmt.Println("length :", len(x))
	fmt.Println(x)
	x = append(x, 45)
	fmt.Println("capacity :", cap(x))
	fmt.Println("length :", len(x))
	fmt.Println(x)
	x = append(x, 46)
	fmt.Println("capacity :", cap(x))
	fmt.Println("length :", len(x))
	fmt.Println(x)
	x = append(x, 47)
	fmt.Println("capacity :", cap(x))
	fmt.Println("length :", len(x))
	fmt.Println(x)

}
