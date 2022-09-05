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

	fmt.Print("Enter a license plate letters: ")
	letters := readInput()
	fmt.Print("Enter a license plate numbers: ")
	numbers, err := strconv.Atoi(readInput())
	if err != nil {
		log.Fatal(err)
	}

	if utf8.RuneCountInString(letters) == 3 && utf8.RuneCountInString(strconv.Itoa(numbers)) == 3 {
		fmt.Println("Old license plate")
	} else if utf8.RuneCountInString(letters) == 3 && utf8.RuneCountInString(strconv.Itoa(numbers)) == 4 {
		fmt.Println("New license plate")
	} else {
		fmt.Println("Not valid license plate")
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}