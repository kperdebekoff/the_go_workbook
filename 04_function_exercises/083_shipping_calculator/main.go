package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter the number of items: ")
	items := readItemsAmount()

	shippingCharge := calcShipping(items)

	fmt.Println("The shipping charge $", shippingCharge)

}

func calcShipping(items float64) float64 {

	const (
		firstItemPrice float64 = 10.95
		nextItemsPrice float64 = 2.95
	)

	shippingCharge := firstItemPrice + (items - 1) * nextItemsPrice

	return shippingCharge

}

func readItemsAmount() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	numberOfItems, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return numberOfItems

}