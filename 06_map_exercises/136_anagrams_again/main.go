package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter some miltiple words: ")
	words1 := readInput()
	fmt.Print("Enter some miltiple words: ")
	words2 := readInput()

	anagrams := isAnagrams(words1, words2)

	fmt.Println("These multiple words are anagrams:", anagrams)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func isAnagrams(words1, words2 string) bool {

	words1 = strings.ToLower(words1)
	words2 = strings.ToLower(words2)

	ignored := []string{".", ",", "!", "?", ":", " "}
	for i := 0; i < len(ignored); i++ {
		for j := 0; j < len(words1); j++ {
			if ignored[i] == string(words1[j]) {
				words1 = strings.Replace(words1, ignored[i], "", -1)
			}
		}
	}
	for i := 0; i < len(ignored); i++ {
		for j := 0; j < len(words2); j++ {
			if ignored[i] == string(words2[j]) {
				words2 = strings.Replace(words2, ignored[i], "", -1)
			}
		}
	}

	if len(words1) == len(words2) {

		uniqChars1 := map[string]bool{}
		for i := 0; i < len(words1); i++ {
			uniqChars1[string(words1[i])] = true
		}
		uniqChars2 := map[string]bool{}
		for i := 0; i < len(words2); i++ {
			uniqChars2[string(words2[i])] = true
		}

		count := 0
		for k1 := range uniqChars1 {
			for k2 := range uniqChars2 {
				if k1 == k2 {
					count += 1
				}
			}
		}

		if count == len(uniqChars1) {
			return true
		} else {
			return false
		}

	} else {

		return false

	}

}