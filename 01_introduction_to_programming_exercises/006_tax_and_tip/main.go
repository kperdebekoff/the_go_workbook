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
		amoutTips = 0.18
		amoutntTax = 0.20
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the cost of a meal: ")
	scanner.Scan()

	amountMoney, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	tip := amountMoney * amoutTips
	tax := amountMoney * amoutntTax
	total := amountMoney + tip + tax

	fmt.Println("Tip:", tip)
	fmt.Println("Tax:", tax)
	fmt.Println("Total:", total)
	
}