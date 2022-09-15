package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter group of words with space separator: ")
	words := readInput()

	pigLatins := convertWordToPigLatin(words)

	fmt.Println(pigLatins)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertWordToPigLatin(s string) []string {

	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		if isBeginY(words[i]) {
			words[i] = yLetterWordToPigLatin(words[i])
		} else if isBeginVowels(words[i]) {
			words[i] = vowelWordToPigLatin(words[i])
		} else {
			words[i] = consonantWordToPigLatin(words[i])
		}
	}

	return words
	
}

func isBeginVowels(s string) bool {

	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	
	isVowel := false
	for i := 0; i < len(s); i++ {
		for _, v := range vowels {
			if s[0] == v {
				isVowel = true
			}
		}
	}

	return isVowel

}

func isBeginY(s string) bool {

	ys := []byte{'y', 'Y'}
	
	isY := false
	for i := 0; i < len(s); i++ {
		for _, v := range ys {
			if s[0] == v {
				isY = true
			}
		}
	}

	return isY

}

func isPunctuationMark(s string) bool {

	punctuation := []byte{',', '.', '!', '?'}

	isPunc := false
	for i := 0; i < len(s); i++ {
		for _, v := range punctuation {
			if s[i] == v {
				isPunc = true
			}
		}
	}

	return isPunc

}

func isBeginUpperCaseLetter(s string) bool {

	if strings.ToUpper(string(s[0])) == string(s[0]) {
		return true
	}
	return false

}

func consonantWordToPigLatin(s string) string {

	w := ""
	if !isBeginVowels(s) && isPunctuationMark(s) && isBeginUpperCaseLetter(s) {
		indVowel := 0
		mid := ""
		end := ""
		for i := 0; i < len(s); i++ {
		if isBeginVowels(string(s[i])) {
			indVowel = i
			break
		}
		}
		for i := 0; i < len(s); i++ {
			if i >= indVowel && !isPunctuationMark(string(s[i])) {
				mid += string(s[i])
			} else if i < indVowel {
				end += strings.ToLower(string(s[i]))
			} else {
				end += "ay" + string(s[i])
			}
		}
		w += strings.Title(mid) + end
	} else if !isBeginVowels(s) && !isPunctuationMark(s) && isBeginUpperCaseLetter(s) {
		indVowel := 0
		mid := ""
		end := ""
		for i := 0; i < len(s); i++ {
		if isBeginVowels(string(s[i])) {
			indVowel = i
			break
		}
		}
		for i := 0; i < len(s); i++ {
			if i >= indVowel {
				mid += string(s[i])
			} else if i < indVowel {
				end += strings.ToLower(string(s[i]))
			}
		}
		w += strings.Title(mid) + end + "ay"
	} else if !isBeginVowels(s) && !isPunctuationMark(s) && !isBeginUpperCaseLetter(s) {
		indVowel := 0
		mid := ""
		end := ""
		for i := 0; i < len(s); i++ {
		if isBeginVowels(string(s[i])) {
			indVowel = i
			break
		}
		}
		for i := 0; i < len(s); i++ {
			if i >= indVowel {
				mid += string(s[i])
			} else if i < indVowel {
				end += strings.ToLower(string(s[i]))
			}
		}
		w += mid + end + "ay"
	} else if !isBeginVowels(s) && isPunctuationMark(s) && !isBeginUpperCaseLetter(s) {
		indVowel := 0
		mid := ""
		end := ""
		for i := 0; i < len(s); i++ {
		if isBeginVowels(string(s[i])) {
			indVowel = i
			break
		}
		}
		for i := 0; i < len(s); i++ {
			if i >= indVowel && !isPunctuationMark(string(s[i])) {
				mid += string(s[i])
			} else if i < indVowel {
				end += strings.ToLower(string(s[i]))
			} else {
				end += "ay" + string(s[i])
			}
		}
		w += mid + end
	} 

	return w

}

func vowelWordToPigLatin(s string) string {

	w := ""
	if isBeginVowels(s) && isPunctuationMark(s) {
		for i := 0; i < len(s); i++ {
			if i < len(s) - 1 {
				w += string(s[i])
			} else {
				w += "way" + string(s[len(s) - 1])
			}
		}
	} else if isBeginVowels(s) && !isPunctuationMark(s) {
		w += s + "way"
	}

	return w

}

func yLetterWordToPigLatin(s string) string {

	w := ""
	if !isBeginVowels(s) && isPunctuationMark(s) {
		for i := 0; i < len(s); i++ {
			if i < len(s) - 1 {
				w += string(s[i])
			} else {
				w += "ay" + string(s[len(s) - 1])
			}
		}
	} else if !isBeginVowels(s) && !isPunctuationMark(s) {
		w += s + "ay"
	}

	return w

}