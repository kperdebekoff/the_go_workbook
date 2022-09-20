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

	fmt.Print(displayPoints(word))

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func displayPoints(word string) int {

	word = strings.ToLower(word)
	wordLetters := []rune(word)

	letterPoints := map[string]int{
		"a": 1, "e": 1, "i": 1, "l": 1, "n": 1, "o": 1, "r": 1, "s": 1, "t": 1, "u": 1,
		"d": 2, "g": 2, "b": 3, "c": 3, "m": 3, "p": 3, "f": 4, "h": 4, "v": 4, "w": 4,
		"y": 4, "k": 5, "j": 8, "x": 8, "q": 10, "z": 10,
	}

	total := 0
	for _, v := range wordLetters {
		for k, val := range letterPoints {
			if string(v) == k {
				fmt.Println(string(v), k, val)
				total += val
			}
		}
	}

	return total

}