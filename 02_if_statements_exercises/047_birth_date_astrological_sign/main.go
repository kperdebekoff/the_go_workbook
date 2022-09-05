package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a month of birth: ")
	month := readInput()
	fmt.Print("Enter a day of birth: ")
	day, err := strconv.Atoi(readInput())
	if err != nil {
		log.Fatal(err)
	}

	if (month == "December" && day >= 22) || (month == "January" && day <= 19) {
		fmt.Println("Your zodiac Capricorn")
	} else if (month == "January" && day >= 20) || (month == "February" && day <= 18) {
		fmt.Println("Your zodiac Aquarius")
	} else if (month == "February" && day >= 19) || (month == "March" && day <= 20) {
		fmt.Println("Your zodiac Pisces")
	} else if (month == "March" && day >= 21) || (month == "April" && day <= 19) {
		fmt.Println("Your zodiac Aries")
	} else if (month == "April" && day >= 20) || (month == "May" && day <= 20) {
		fmt.Println("Your zodiac Taurus")
	} else if (month == "May" && day >= 21) || (month == "June" && day <= 20) {
		fmt.Println("Your zodiac Gemini")
	} else if (month == "June" && day >= 21) || (month == "July" && day <= 22) {
		fmt.Println("Your zodiac Cancer")
	} else if (month == "July" && day >= 23) || (month == "August" && day <= 22) {
		fmt.Println("Your zodiac Leo")
	} else if (month == "August" && day >= 23) || (month == "September" && day <= 22) {
		fmt.Println("Your zodiac Vigro")
	} else if (month == "September" && day >= 23) || (month == "October" && day <= 22) {
		fmt.Println("Your zodiac Libra")
	} else if (month == "October" && day >= 23) || (month == "November" && day <= 21) {
		fmt.Println("Your zodiac Scorpio")
	} else if (month == "November" && day >= 22) || (month == "December" && day <= 21) {
		fmt.Println("Your zodiac Sagittarius")
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}