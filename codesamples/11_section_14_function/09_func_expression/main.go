package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("This is first func expression")
	}
	a()
	b := func(x int) {
		fmt.Println("I watched on video persistently ", x)

	}
	b(45)
}
