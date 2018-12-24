package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `P@ssw0rd`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(s, " ------> ", string(bs))
	
	loginPass:=`P@ssw0rd`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err!= nil {
		fmt.Println("You cant log in")
		return
	}
	fmt.Println("You are logged in")
}
