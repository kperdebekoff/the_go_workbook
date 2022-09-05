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

	fmt.Print("Enter a first side length: ")
	s1 := readInput()
	fmt.Print("Enter a second side length: ")
	s2 := readInput()
	fmt.Print("Enter a third side length: ")
	s3 := readInput()

	s := (s1 + s2 + s3) / 2
	area := math.Sqrt(s * (s - s1) * (s - s2) * (s - s3))

	fmt.Println("Triangle area:", area)

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