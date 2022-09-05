package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter your weight in kg: ")
	weight := readInput()
	fmt.Print("Enter your height in m: ")
	height := readInput()

	bmi := weight / (height * height)

	fmt.Print("Your BMI (body mass index): ", bmi)
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