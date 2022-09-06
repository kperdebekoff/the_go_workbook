package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter encoded or decoded text: ")
	text := readInputText()
	fmt.Print("Enter positive step for encode and negative for decode: ")
	shift := readInputShift()

	message := encodeDecodeMessage(text, shift)
	fmt.Println(string(message))

}

func readInputText() []rune {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return []rune(scan.Text())

}

func readInputShift() int {

	shift, err := strconv.Atoi(string(readInputText()))
	if err != nil {
		log.Fatal(err)
	}
	return shift

}

func encodeDecodeMessage(text []rune, shift int) []rune {

	for i := 0; i < len(text); i++ {

		if text[i] >= 65 && text[i] <= 90 && text[i] + rune(shift) > 90 {
			text[i] = text[i] + rune(shift) - 26
		} else if text[i] >= 65 && text[i] <= 90 && text[i] + rune(shift) < 90 {
			text[i] = text[i] + rune(shift)
		} else if text[i] >= 97 && text[i] <= 122 && text[i] + rune(shift) > 122 {
			text[i] = text[i] + rune(shift) - 26
		} else if text[i] >= 97 && text[i] <= 122 && text[i] + rune(shift) < 122 {
			text[i] = text[i] + rune(shift)
		} else {
			text[i] = text[i] + rune(shift)
		}
	}
	return text

}