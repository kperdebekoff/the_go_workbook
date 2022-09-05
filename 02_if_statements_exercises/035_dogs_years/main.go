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
	fmt.Println("Enter a number of years to the convert from human years to dog years:")
	scan.Scan()
	humanYears, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	firstTwoYears := 2 * 10.5

	if humanYears < 0 {
		fmt.Println("Please enter a positive number of years")
	} else if humanYears > 2 {
		dogYears := (humanYears - 2) * 4 + firstTwoYears
		fmt.Println(humanYears, "human years is equivalent to", dogYears, "dog years")
	} else if humanYears == 1 {
		fmt.Println(humanYears, "human years is equivalent to", firstTwoYears / 2, "dog years")
	} else {
		fmt.Println(humanYears, "human years is equivalent to", firstTwoYears, "dog years")
	}
	
}