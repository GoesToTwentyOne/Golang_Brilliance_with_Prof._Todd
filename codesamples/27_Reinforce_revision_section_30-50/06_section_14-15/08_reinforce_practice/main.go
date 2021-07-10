package main

import "fmt"

func main() {
	fmt.Println(goes()())
	// g:=goes()

}
func goes() func() int {
	return func() int {
		return 42

	}
}
