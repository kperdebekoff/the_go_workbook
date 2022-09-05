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

	fmt.Print("Enter radius of a cylinder: ")
	radius := readInput()
	fmt.Print("Enter height of a cylinder: ")
	height := readInput()

	volume := math.Pi * math.Pow(radius, 2) * height

	fmt.Println("Volume of a cylinder:", math.Floor(volume * 10) / 10)

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