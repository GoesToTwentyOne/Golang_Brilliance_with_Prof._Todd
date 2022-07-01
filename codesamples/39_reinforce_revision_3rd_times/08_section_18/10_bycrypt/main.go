package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `Nehagolang123@#`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	//password match
	login_pass := `Nehagolang123@#`
	err = bcrypt.CompareHashAndPassword(bs, []byte(login_pass))
	if err != nil {
		fmt.Println(err)
		fmt.Println("login unsuccessful")
	} else {
		fmt.Println("login successful")
	}

}
