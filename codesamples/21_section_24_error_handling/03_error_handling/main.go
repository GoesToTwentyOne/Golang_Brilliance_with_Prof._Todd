package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("read.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// fmt.Println(f)
	read := strings.NewReader("Hello this is Nihad")
	w, err := io.Copy(f, read)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w)
}
