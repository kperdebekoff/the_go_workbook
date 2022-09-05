package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a year: ")
	scan.Scan()
	year, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

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
		fmt.Println(year, " year - leap years")
	} else {
		fmt.Println(year, " year - not leap years")
	}

}