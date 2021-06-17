package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `Togent123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))
	err = bcrypt.CompareHashAndPassword(bs, []byte(`Togent123`))
	if err != nil {
		fmt.Println(err)
		fmt.Println("your password don't match")
	} else {
		fmt.Println("your password  match")

	}

}
