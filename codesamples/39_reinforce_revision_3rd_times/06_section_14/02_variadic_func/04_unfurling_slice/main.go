package main

import "fmt"

func main() {
	xs_T("Nihad")
	xs_TS("Bangladesh", "FG", "UG", "HG")
	xs := []string{"first", "Second", "Third"}
	xs_TS("Hello", xs...)
}

func xs_T(s string, a ...string) {
	fmt.Println(s)
	fmt.Println(a)
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", a)

}
func xs_TS(s string, a ...string) {
	fmt.Println(s)
	fmt.Println(a)
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", a)

}
