package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("nihad.txt")
	if err != nil {
		fmt.Println(err)
	}
	// s := strings.NewReader("Hello grpc coding")
	// _, err = io.Copy(f, s)
	_, err = f.WriteString("Hello from write")
	if err != nil {
		fmt.Println(err)
	}

	f, err = os.Open("nihad.go")
	if err != nil {
		fmt.Println(err)
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
