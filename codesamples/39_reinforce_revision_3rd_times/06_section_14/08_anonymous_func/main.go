package main

import "fmt"

func main() {
	func(age int, y, z string) {
		fmt.Println("age is :", age)
		fmt.Printf("name is %s %s\n", y, z)

	}(21, "Nihad", "Hossain")
}
