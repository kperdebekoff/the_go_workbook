package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter number of day from 1 to 31: ")
	day := readInput()
	fmt.Print("Enter number of month from 1 to 12: ")
	month := readInput()
	fmt.Print("Enter a year: ")
	year := readInput()

	magicDate := isMagicDate(day, month, year)

	fmt.Println(magicDate)

}

func isMagicDate(day, month, year int) bool {

	isMagic := false
	if month * day == year % 100 {
		isMagic = true
	}
	return isMagic

}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	n, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	return n

}