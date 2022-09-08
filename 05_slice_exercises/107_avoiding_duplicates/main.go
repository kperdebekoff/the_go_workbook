package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Enter a word: ")
	word := readInput()
	if word == "" {
		log.Fatal("Empty line")
	}

	wordsSlice := storeDataInSlice(word)
	for _, v := range wordsSlice {
		fmt.Println(v)
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func storeDataInSlice(word string) []string {
	
	var words []string

	contains := func(w []string, s string) bool {
		for _, v := range w {
			if v == s {
				return true
			}
		}
		return false
	}

	for word != "" {
		if !contains(words, word) {
			words = append(words, word)
		}
		fmt.Print("Enter a word (blank line to quit): ")
		word = readInput()
	}

	return words

}
