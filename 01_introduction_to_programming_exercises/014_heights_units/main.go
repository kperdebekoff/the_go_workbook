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
		inchesInFoot = 12
		cmInInches = 2.54
	)
	
	fmt.Print("Enter number of feets in your height: ")
	feets := readInput()

	fmt.Print("Enter number of inches in your height: ")
	inches := readInput()

	cm := feets * (inchesInFoot * cmInInches) + inches * cmInInches

	fmt.Println("Number of centimeters in your height", math.Floor(cm * 100) / 100)

}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	height, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	return height

}