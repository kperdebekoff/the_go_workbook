package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a letter of the alphabet in lowercase: ")
	scan.Scan()
	letter := scan.Text()

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Print("The entered letter is a vowel")
	case "y":
		fmt.Print("The entered letter sometimes is a vowel and sometimes is a consonant")
	default:
		fmt.Print("The entered letter is a consonant")
	}

}