package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter some words: ")
	text := readInput()
	uniqChars := findAmountUniqChar(text)

	fmt.Print(uniqChars)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func findAmountUniqChar(text string) int {

	text = strings.ToLower(text)
	uniqChar := map[string]bool{}

	for i := 0; i < len(text); i++ {
		uniqChar[string(text[i])] = true
	}

	return len(uniqChar)

}