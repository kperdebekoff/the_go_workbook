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
	fmt.Print("Enter a first integer number: ")
	first := readInput()
	fmt.Print("Enter a second integer number: ")
	second := readInput()
	fmt.Print("Enter a third integer number: ")
	third := readInput()

	maxVal := math.Max(math.Max(first, second), third)
	minVal := math.Min(math.Min(first, second), third)
	midVal := (first + second + third) - (maxVal + minVal)

	fmt.Println("From smallest to largest:", minVal, midVal, maxVal)

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