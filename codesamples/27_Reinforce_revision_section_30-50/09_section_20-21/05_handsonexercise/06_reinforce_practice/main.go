package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("ARCH ", runtime.GOARCH)
	fmt.Println("OS", runtime.GOOS)
}
