package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter group of words in lowercase with space separator: ")
	words := readInput()

	fmt.Println(convertToPigLatin(words))

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertToPigLatin(w string) []string {

	sliceOfWords := strings.Split(w, " ")

	for i := 0; i < len(sliceOfWords); i++ {

		if isBeginVowels(sliceOfWords[i]) {
			sliceOfWords[i] += "way"
		} else {
			w := ""
			symb := []rune(sliceOfWords[i])

			if symb[0] != 'y' {
				for i := 0; i < len(symb); i++ {
					if i > 0 {
						w += string(symb[i])
					}
				}
				w += string(symb[0]) + "ay"
				sliceOfWords[i] = w
			} else {
				sliceOfWords[i] += "ay"
			}
		}

	}

	return sliceOfWords

}

func isBeginVowels(word string) bool {

	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	isVowel := false
	for _, v := range vowels {
		if v == word[0] {
			isVowel = true
		}
	}
	return isVowel

}