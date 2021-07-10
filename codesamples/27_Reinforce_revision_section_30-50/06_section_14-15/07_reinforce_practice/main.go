package main

import "fmt"

func main() {
	f := func(s string, i int) string {
		return fmt.Sprintf("Name :%s Age : %d", s, i)

	}
	fmt.Printf("%T\n", f)
	fmt.Println(f("Nihad", 21))

}
