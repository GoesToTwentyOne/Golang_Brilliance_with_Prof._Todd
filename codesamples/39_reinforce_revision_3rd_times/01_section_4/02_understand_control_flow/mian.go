package main

import "fmt"

func main() {
	fmt.Println("Hello this is line 4")
	going()
	fmt.Println("after funciton calling")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
	fmt.Println("executing the program")

}
func going() {
	fmt.Println("Hello this is nihad practice code")
}

//Control flow
//01-sequenctial
//02-loop,iterative
//03-conditional
