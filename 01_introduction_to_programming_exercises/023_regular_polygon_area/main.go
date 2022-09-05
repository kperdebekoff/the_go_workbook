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

	fmt.Print("Enter the length of a polygon side: ")
	s := readInput()
	fmt.Print("Enter the number of a polygon sides: ")
	n := readInput()

	area := (n * math.Pow(s, 2)) / (4 * math.Tan(math.Pi / n))

	fmt.Println("Area of a regular polygon:", area)

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