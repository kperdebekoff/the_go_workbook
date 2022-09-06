package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode/utf8"
)

func main() {

	fmt.Print("Enter a password: ")
	password := readInput()
	goodPassword := checkPassword(password)

	fmt.Println("Password passed status:", goodPassword)

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

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}