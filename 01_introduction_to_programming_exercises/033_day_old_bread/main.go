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

	const breadPrice = 3.49

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number of loaves of day old breads: ")
	scan.Scan()
	num, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	discountPrice := (breadPrice - breadPrice * 0.6)
	oldBreadPrice := num * discountPrice

	fmt.Println("Regular price for bread:", breadPrice)
	fmt.Println("Discount price for old bread:", math.Floor(discountPrice * 100) / 100)
	fmt.Println("Total price for old bread:", math.Floor(oldBreadPrice * 100) / 100)

}