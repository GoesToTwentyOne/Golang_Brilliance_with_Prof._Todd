package main

import "fmt"

func main() {
	s := "Nihad Hossain"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])

	}
	for index, v := range s {
		fmt.Printf("%#U \t %d", v, index)
	}
}
