package main

import (
	"fmt"
)

func main() {
	var aa, bb, cc string
	fmt.Println("Enter one string :")
	_, err := fmt.Scanf("%s\n", &aa)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter one string :")
	_, err = fmt.Scanf("%s\n", &bb)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter one string :")
	_, err = fmt.Scanf("%s\n", &cc)
	if err != nil {
		panic(err)
	}
	fmt.Println(aa, bb, cc)

}
