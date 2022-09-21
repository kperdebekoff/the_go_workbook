package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter some text to convert to key press: ")
	text := readInput()

	fmt.Println(convertStrToButtonClick(text))
}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertStrToButtonClick(text string) string {

	symbToKeyPress := map[string]string{
		".": "1",
		",": "11",
		"?": "111",
		"!": "1111",
		":": "11111",
		"a": "2",
		"b": "22",
		"c": "222",
		"d": "3",
		"e": "33",
		"f": "333",
		"g": "4",
		"h": "44",
		"i": "444",
		"j": "5",
		"k": "55",
		"l": "555",
		"m": "6",
		"n": "66",
		"o": "666",
		"p": "7",
		"q": "77",
		"r": "777",
		"s": "7777",
		"t": "8",
		"u": "88",
		"v": "888",
		"w": "9",
		"x": "99",
		"y": "999",
		"z": "9999",
		" ": "0",
	}

	text = strings.ToLower(text)
	arrRune := []rune(text)
	keyPress := ""
	for i := 0; i < len(arrRune); i++ {
		if arrRune[i] >= 97 || arrRune[i] <= 122 {
			keyPress += symbToKeyPress[string(arrRune[i])]
		} else if arrRune[i] == 32 || arrRune[i] == 58 {
			keyPress += symbToKeyPress[string(arrRune[i])]
		} else if arrRune[i] == 33 || arrRune[i] == 44 || arrRune[i] == 46 || arrRune[i] == 63 {
			keyPress += symbToKeyPress[string(arrRune[i])]
		}
	}
	
	return keyPress
	
}