package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter first word: ")
	word1 := readInput()
	fmt.Print("Enter second word: ")
	word2 := readInput()

	anagram := isAnagrams(word1, word2)

	fmt.Println("Words are anagrams:", anagram)

}

func readInput() string {
	
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func isAnagrams(word1, word2 string) bool {

	if len(word1) == len(word2) {

		uniqChar1 := map[string]string{}
		uniqChar2 := map[string]string{}

		for i := 0; i < len(word1); i++ {
			uniqChar1[string(word1[i])] = string(word1[i])
		}
	
		for i := 0; i < len(word2); i++ {
			uniqChar2[string(word2[i])] = string(word2[i])
		}

		count := 0
		for k1 := range uniqChar1 {
			for k2 := range uniqChar2 {
				if k1 == k2 {
					count += 1
				}
			}
		}
		return len(uniqChar1) == count

	} else {
		return false
	}

}