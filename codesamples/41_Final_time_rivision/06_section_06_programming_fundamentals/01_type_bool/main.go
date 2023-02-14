package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(2 == 3)
	fmt.Println(2 <= 6)
	fmt.Println(4 >= 67)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
