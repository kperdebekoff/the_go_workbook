package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	const (
		oneLitersDeposit = 0.10
		moreOneLitersDeposit = 0.25
	)
	
	fmt.Println("Enter amount of 1 liter or less bottle:")
	smallBottle := convertstrToFloat(readEnteredText())

	fmt.Println("Enter amount of more than 1 liter bottle:")
	largeBottle := convertstrToFloat(readEnteredText())

	money := smallBottle * oneLitersDeposit + largeBottle * moreOneLitersDeposit

	fmt.Println("Refund $", money)
}

func readEnteredText() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

func convertstrToFloat(a string) float64 {

	amount, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatal(err)
	}
	return amount

}