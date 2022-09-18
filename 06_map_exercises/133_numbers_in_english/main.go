package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a positive integer number from 1 to 999: ")
	number := readInput()

	word := convertNumToWord(number)

	fmt.Println(word)

}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return num

}

func convertNumToWord(number int) string {

	oneToNine := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven",
		8: "eight", 9: "nine",
	}

	tenToNineteen := map[int]string{
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
	}

	twentyToNinety := map[int]string{
		20: "twenty", 30: "thirty", 40: "fourty", 50: "fifty", 60: "sixty", 70: "seventy",
		80: "eighty", 90: "ninety",
	}

	oneHundredToNineHundred := map[int]string{}
	for i := 1; i <= 9; i++ {
		oneHundredToNineHundred[i * 100] = oneToNine[i] + " hundred"
	}

	// from 20 to 99 add to twentyToNinety
	for i := 20; i <= 90; i += 10 {
		for j := 1; j <= 9; j++ {
			twentyToNinety[i + j] = twentyToNinety[i] + " " + oneToNine[j]
		}
	}

	// from 100 to 999 add to oneHundredToNineHundred map
	for i := 100; i <= 900; i += 100 {
		for j := 1; j <= 9; j++ {
			oneHundredToNineHundred[i + j] = oneHundredToNineHundred[i] + " " + oneToNine[j]
		}
		for k := 10; k <= 19; k++ {
			oneHundredToNineHundred[i + k] = oneHundredToNineHundred[i] + " " + tenToNineteen[k]
		}
		for l := 20; l <= 99; l++ {
			oneHundredToNineHundred[i + l] = oneHundredToNineHundred[i] + " " + twentyToNinety[l]
		}
	}

	word := ""
	if number < 10 {
		word = oneToNine[number]
	} else if number >= 10 && number < 20 {
		word = tenToNineteen[number]
	} else if number >= 20 && number < 100 {
		word = twentyToNinety[number]
	} else if number >= 100 && number < 999 {
		word = oneHundredToNineHundred[number]
	} else {
		word = "Please enter an integer number from 1 to 999"
	}

	return word
	
}