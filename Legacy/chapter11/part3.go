package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// bcrypt
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	loginPwd := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Logged in")

}
