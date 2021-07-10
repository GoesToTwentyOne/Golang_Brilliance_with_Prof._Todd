package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i, value := range arr {
		fmt.Println(i, "----- ", value)
	}
	fmt.Printf("%T", arr)

}
