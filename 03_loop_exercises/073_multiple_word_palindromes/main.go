package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter a phrases: ")
	phrases := readInput()
	runePhrases := []rune(strings.ReplaceAll(strings.ToLower(phrases), " ", ""))
	isPalindrome := true

	for i := 0; i < len(runePhrases); i++ {
		if runePhrases[i] != runePhrases[len(runePhrases) - i - 1] {
			isPalindrome = false
		}
	}
	
	if isPalindrome {
		fmt.Println("This phrases is palindrom:", phrases)
	} else {
		fmt.Println("This phrases is NOT palindrom:", phrases)
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}