package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	totalCost := 0.0
	for {
		fmt.Print("Enter the age of visitors: ")
		ages := readInput()
		if ages == -1.0 {
			break
		}
		cost := calcTicketPrice(ages)
		totalCost += cost
	}

	fmt.Println("Total cost for admissions: $", totalCost)

}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	age, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		return -1.0
	}
	return age

}

func calcTicketPrice(age float64) float64 {

	price := 0.00
	if age <= 2 {
		price = 0.00
	} else if age >= 3 && age <= 12 {
		price = 14.00
	} else if age >= 13 && age <= 65 {
		price = 23.00
	} else {
		price = 18.00
	}
	return price

}