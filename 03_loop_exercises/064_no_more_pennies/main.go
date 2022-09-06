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

	const (
		penniesPerNickel = 5
		nickel = 0.05
	)
	var totalCost float64 = 0
	var cash float64 = 0

	for {
		fmt.Print("Enter a price: ")
		price := readInput()
		if price == "" {
			break
		}

		cost, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal(err)
		}
		totalCost += cost
	}

	totalCost = math.Floor(totalCost * 100) / 100
	roundIndicator := int(totalCost * 100) % penniesPerNickel
	if roundIndicator > penniesPerNickel / 2 {
		cash = totalCost - float64(roundIndicator) / 100 + nickel
	} else {
		cash = totalCost - float64(roundIndicator) / 100
	}

	fmt.Println("Total cost:", totalCost)
	fmt.Println("Pay with cash:", math.Floor(cash * 100) / 100)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}