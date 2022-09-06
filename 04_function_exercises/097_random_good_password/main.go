package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"time"
	"unicode/utf8"
)

func main() {

	numAttempts := countAttempts()
	fmt.Println("Number of attempts for generate good password:", numAttempts + 1)
	
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

func checkPassword(p string) bool {

	isGood := false

	containSymb := func(s string) bool {
		symb, err := regexp.MatchString(s, p)
	if err != nil {
		log.Fatal(err)
	}
	return symb
	}
	length := utf8.RuneCountInString(p) >= 8
	
	if length && containSymb(`[A-Z]+`) && containSymb(`[a-z]+`) && containSymb(`[0-9]+`) {
		isGood = true
	}
	return isGood

}

func countAttempts() int {

	pass := generatePassword()
	count := 0
	for !checkPassword(pass) {
		count += 1
		pass = generatePassword()
	}
	return count

}