package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter some water mass in grams or ml: ")
	mass := readInput()
	fmt.Print("Enter the temperature change in degrees Celcius: ")
	temperature := readInput()

	q := mass * 4.186 * temperature
	fmt.Println("Total amount of energy that must be added or removed:", q, "Joules")

	cost := 8.9 * q * 2.77778e-7
	fmt.Println("Costs of heating water:", cost, "cents")

}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	data, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return data

}