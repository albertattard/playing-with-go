package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func runPasswordsDemo() {
	plaintextPassword := "This is my long and secure password"
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 13)

	if e != nil {
		fmt.Printf("Failed to encode password %v", e)
		return
	}

	fmt.Printf("Plain Text: %v\n", plaintextPassword)
	fmt.Printf("Encoded:    %v\n", string(hashedPassword))

	challenge1 := "This is my long and secure password"
	passwordMatchError1 := bcrypt.CompareHashAndPassword(hashedPassword, []byte(challenge1))
	if passwordMatchError1 == nil {
		fmt.Println("Password match!!")
	} else {
		fmt.Println("Password does not match!!  Please try again")
	}

	challenge2 := "This is a different password"
	passwordMatchError2 := bcrypt.CompareHashAndPassword(hashedPassword, []byte(challenge2))
	if passwordMatchError2 == nil {
		fmt.Println("Password match!!")
	} else {
		fmt.Println("Password does not match!!  Please try again")
	}
}
