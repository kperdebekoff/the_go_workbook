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
	
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter temperature in Celsius degree: ")
	scan.Scan()
	tempInCelc, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	tempInFah := 33.8 * tempInCelc
	tempInKel := 273.15 + tempInCelc

	fmt.Println(tempInCelc, "in Celsius degree is", math.Round(tempInFah), "Fahrenheit degrees")
	fmt.Println(tempInCelc, "in Celsius degree is", math.Round(tempInKel), "Kelvin degrees")

}