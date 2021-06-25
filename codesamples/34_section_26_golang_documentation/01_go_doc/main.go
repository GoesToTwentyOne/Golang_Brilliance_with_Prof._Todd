package main

import "fmt"

func main() {
	Goes()
	doc("Goes To Twenty One")

}
func Goes() {
	fmt.Println("goes to twenty one")
}
func doc(s string) {
	fmt.Printf("%sis my github user name \n", s)
}
