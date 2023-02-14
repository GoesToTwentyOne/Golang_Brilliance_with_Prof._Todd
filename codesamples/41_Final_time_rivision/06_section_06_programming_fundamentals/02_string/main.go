package main

import (
	"fmt"
)

func main() {
	s := "Md. Niahd Hossain"
	fmt.Println(s)
	fmt.Printf("\n%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("\n%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	for i, v := range s {
		fmt.Printf("\nindex is %d value of the slice %v", i, v)

	}
}
