package main

import "fmt"

func main() {
	n := `"My name is \'MD.Nihad Hossain\'"`
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	//slice of byte
	bs := []byte(n)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
	for i, v := range n {
		fmt.Printf("index is %d decimal value is %d hexadecimal value is %#x  unicode %#U\n", i, v, v, v)
	}
}
