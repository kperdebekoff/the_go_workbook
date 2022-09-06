package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a number of month (from 1 to 12): ")
	month := readInput()
	fmt.Print("Enter a year: ")
	year := readInput()

	days := displayDaysInMonths(month, year)

	fmt.Println("year:", year, "month:", month, "days:", days)
	
}

func displayDaysInMonths(month int, year int) int {

	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 2:
		if year % 4 == 0 || (year % 100 == 0 && year % 400 == 0) {
			day = 29
		} else {
			day = 28
		}
	case 4, 6, 9, 11:
		day = 30
	}

	return day

}

func readInput() int {
	
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	d, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return d

}