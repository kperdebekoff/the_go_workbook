package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	pswrd := generatePassword()
	fmt.Println("Generated password:", pswrd)

}

func generatePassword() string {

	password := ""
	min := 33
	max := 126
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randChar := rand.Intn(max - min + 1) + min
		password += string(randChar)
	}
	return password
	
}