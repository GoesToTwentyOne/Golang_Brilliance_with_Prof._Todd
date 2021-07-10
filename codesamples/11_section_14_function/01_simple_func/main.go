package main

import (
	"fmt"
)

func main() {
	goes()
	a := bee("Nihad")
	b, c := cube("Scotiya", 21)
	fmt.Println(a)
	fmt.Println(b, c)

}
func goes() {
	fmt.Println("Hello from goes")
}
func bee(s string) string {
	return fmt.Sprintf(`Hi, My name is %s`, s)

}
func cube(s string, p int) (int, string) {
	return p, fmt.Sprintf("\n%sLaaa la baby %v I'm on your heart", s, p)

}
