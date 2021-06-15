package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is anonymous func")
	}()
	func(a int) {
		fmt.Printf(`life is type when you failed %v times you think that is simple "Programmer"`, a)

	}(75)
}
