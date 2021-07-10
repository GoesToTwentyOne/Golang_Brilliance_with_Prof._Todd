package main

import "fmt"

func main() {
	i := 1
	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("done")

}
