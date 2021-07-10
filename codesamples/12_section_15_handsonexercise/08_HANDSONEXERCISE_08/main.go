package main

import "fmt"

func main() {
	// rF := reFunc()
	// fmt.Println(rF())
	fmt.Println(reFunc()())

}
func reFunc() func() string {
	return func() string {
		return "Hello from ruturning a func"
	}

}
