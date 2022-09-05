package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter a position letter: ")
	positionLetter := readInput()
	fmt.Print("Enter a position number: ")
	positionNumber, err := strconv.Atoi(readInput())
	if err != nil {
		log.Fatal(err)
	}

	pos := ""
	switch positionLetter {
	case "a", "c", "e", "g":
		pos = "a"
	case "b", "d", "f", "h":
		pos = "b"
	}

	num := 0
	switch positionNumber {
	case 1, 3, 5, 7:
		num = 1
	case 2, 4, 6, 8:
		num = 2
	}

	if pos == "a" && num == 1 {
		fmt.Println("Black square")
	} else if pos == "a" && num == 2 {
		fmt.Println("White square")
	} else if pos == "b" && num == 1 {
		fmt.Println("White square")
	} else if pos == "b" && num == 2 {
		fmt.Println("Black square")
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}