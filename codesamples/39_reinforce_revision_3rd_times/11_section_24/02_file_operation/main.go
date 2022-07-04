package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//write
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := strings.NewReader("Nihad Hossain")
	io.Copy(f, r)
	//read
	f, err = os.Open("name.txt")
	if err != nil {
		fmt.Println(err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
