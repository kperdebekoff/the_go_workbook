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
		if !isBeginVowels(sliceOfWords[i]) {
			sliceOfWords[i] = consonantWordToPigLatin(sliceOfWords[i])
		} else {
			sliceOfWords[i] = vowelWordToPigLatin(sliceOfWords[i])
		}
	}
	return sliceOfWords
}

func isBeginVowels(word string) bool {

	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	isVowel := false
	for _, v := range vowels {
		if word[0] == v {
			isVowel = true
		}
	}
	return isVowel

}

func consonantWordToPigLatin(word string) string {

	w := ""
if word[0] != 'y' {

	ind := 0
	end := ""
	for i := 0; i < len(word); i++ {
		if isBeginVowels(string(word[i])) {
			ind = i
			break
		}
	}

	for i := 0; i < len(word); i++ {
		if i >= ind {
			w += string(word[i])
		} else {
			end += string(word[i])
		}
	}

	w += end + "ay"

} else {

	for i := 0; i < len(word); i++ {
		w += string(word[i])
	}

	w += "ay"

}

	return w

}

func vowelWordToPigLatin(word string) string {

	w := ""
	for i := 0; i < len(word); i++ {
		w += string(word[i])
	}
	w += "way"

	return w

}