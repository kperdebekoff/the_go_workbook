package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	
	fmt.Print("Enter a string of characters: ")
	text := readUserInput()
	fmt.Print("Enter width of the terminal in characters: ")
	terminalWidth, err := strconv.Atoi(readUserInput())
	if err != nil {
		log.Fatal(err)
	}

	centeredString := displayCenteredText(text, terminalWidth)

	fmt.Println(centeredString)

}

func readUserInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func displayCenteredText(s string, w int) string {

	if utf8.RuneCountInString(s) >= w {
		return s
	} else {
		space := func(s string, w int) string {
			amountSpace := ""
			for i := 0; i < (w - utf8.RuneCountInString(s)) / 2; i++ {
				amountSpace += " "
			}
			return amountSpace + s + amountSpace
		}

		return space(s, w)
	}

}