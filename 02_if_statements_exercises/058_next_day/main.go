package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a year: ")
	year := readInput()
	fmt.Print("Enter a month: ")
	month := readInput()
	fmt.Print("Enter a day: ")
	day := readInput()

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day == 31 && month != 12 {
			month += 1
			day = 1
		} else if day == 31 && month == 12 {
			year += 1
			month = 1
			day = 1
		} else {
			day += 1
		}

	case 4, 6, 9, 11:
		if day == 30 {
			month += 1
			day = 1
		} else {
			day += 1
		}

	case 2:
		isLeapYear := false
	
		if year % 400 == 0 {
			isLeapYear = true
		} else if year % 100 == 0 {
			isLeapYear = false
		} else if year % 4 == 0 {
			isLeapYear = true
		} else {
			isLeapYear = false
		}

		if isLeapYear {
			if day == 29 {
				month += 1
				day = 1
			} else {
				day += 1
			}
		} else {
			if day == 28 {
				month += 1
				day = 1
			} else {
				day += 1
			}
		}
	}
	
	if month < 10 && day < 10 {
		fmt.Println("Date:", year, "-", "0" + strconv.Itoa(month), "-", "0" + strconv.Itoa(day))
	} else if month < 10 && day > 10 {
		fmt.Println("Date:", year, "-", "0" + strconv.Itoa(month), "-", day)
	} else if month > 10 && day < 10 {
		fmt.Println("Date:", year, "-", month, "-", "0" + strconv.Itoa(day))
	} else {
		fmt.Println("Date:", year, "-", month, "-", day)
	}
	
}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	date, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return date

}