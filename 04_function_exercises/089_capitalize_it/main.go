package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	
	fmt.Print("Enter some text to capitalize: ")
	s := readString()

	capitalize := correctText(s)

	fmt.Println(capitalize)

}

func correctText(s string) string {

	text := []rune(s)
	// For ".!? letter" after ".", "!", "?" --> ".!? Letter"
	re, err := regexp.Compile(`^\w|[\.\!\?] \w`)
	if err != nil {
		log.Fatal(err)
	}
	res := re.FindAllStringIndex(s, -1)
	for i := 0; i < len(res); i++ {
		for j := 1; j < len(res[i]); j++ {
			text[(res[i][j] - 1)] = text[(res[i][j] - 1)] - 32
		}
	}

	// For "i" letter --> "I" letter
	iLetter := func(t []rune) []rune {
		re, err := regexp.Compile(`[^.]\si|\bi\b`)
	if err != nil {
		log.Fatal(err)
	}
	res := re.FindAllStringIndex(s, -1)
	for i := 0; i < len(res); i++ {
		for j := 1; j < len(res[i]); j++ {
			t[(res[i][j] - 1)] = t[(res[i][j] - 1)] - 32
		}
	}
	return t
	}
	iLetter(text)
	
	return string(text)

}

func readString() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}