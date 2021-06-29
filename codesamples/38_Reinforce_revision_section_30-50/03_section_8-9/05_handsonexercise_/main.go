package main

import "fmt"

func main() {
	x := 10
	for {
		if x > 100 {
			break
		}
		if x%4 == 0 {
			fmt.Println("divideble by 4 : ", x)
		}
		x++

	}

}
