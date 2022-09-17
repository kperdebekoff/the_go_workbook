package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter some message to convert Morse Code: ")
	message := readInput()

	morseCodeMessage := convertMessageToMorseCode(message)

	fmt.Println(morseCodeMessage)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertMessageToMorseCode(message string) string {

	morseCodeMessage := ""

	symbInMorseCode := map[string]string{
		"a": ".-", "b": "-...", "c": "-.-.", "d": "-..", "e": ".", "f": "..-.",
		"g": "--.", "h": "....", "i": "..", "j": ".---", "k": "-.-", "l": ".-..",
		"m": "--", "n": "-.", "o": "---", "p": ".--.", "q": "--.-", "r": ".-.",
		"s": "...", "t": "-", "u": "..-", "v": "...-", "w": ".--", "x": "-..-",
		"y": "-.---", "z": "--..", "0": "-----", "1": ".----", "2": "..---",
		"3": "...--", "4": "....-", "5": ".....", "6": "-....", "7": "--...",
		"8": "---..", "9": "----.",
	}

	message = strings.ToLower(message)
	arrRune := []rune(message)
	for i := 0; i < len(arrRune); i++ {
		if arrRune[i] >= 97 && arrRune[i] <= 122 {
			morseCodeMessage += symbInMorseCode[string(arrRune[i])] + " "
		} else if arrRune[i] >= 48 && arrRune[i] <= 57 {
			morseCodeMessage += symbInMorseCode[string(arrRune[i])] + " "
		} else {
			morseCodeMessage += ""
		}
	}

	return morseCodeMessage
	
}