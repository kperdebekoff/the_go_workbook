package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	const interestPerYear = 0.04

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter amount of money on your deposites:")
	scanner.Scan()
	money := convertStrToFloat(scanner.Text())
	
	saveFirstYear := money * interestPerYear + money
	saveSecondYear := (money + saveFirstYear) * interestPerYear + money
	saveThirdYear := (money + saveSecondYear) * interestPerYear + money

	fmt.Println("Savings after 1 year:", (math.Floor(saveFirstYear * 100) / 100))
	fmt.Println("Savings after 2 years:", (math.Floor(saveSecondYear * 100) / 100))
	fmt.Println("Savings after 3 years:", (math.Floor(saveThirdYear * 100) / 100))

}

func convertStrToFloat(s string) float64 {

	amount, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return amount
}