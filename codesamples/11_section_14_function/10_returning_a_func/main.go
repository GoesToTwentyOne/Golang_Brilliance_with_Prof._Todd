package main

import "fmt"

func main() {
	fmt.Println(goes()())

}
func goes() func() string {
	return func() string {
		return "I love  my country"
	}

}
