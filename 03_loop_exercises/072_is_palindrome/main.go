package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter a word: ")
	word := readInput()
	text := []rune(strings.ToLower(word))
	
	isPalindrome := true
	for i := 0; i < len(text); i++ {
		if text[i] != text[len(text) - i - 1] {
			isPalindrome = false
		}
	}

	if isPalindrome {
		fmt.Println("This text is palindrom:", word)
	} else {
		fmt.Println("This text is not palindrome:", word)
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}