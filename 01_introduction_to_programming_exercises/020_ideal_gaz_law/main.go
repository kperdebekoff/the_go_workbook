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

	const R = 8.314

	fmt.Print("Enter gas pessure in Pascal: ")
	p := readInput()
	fmt.Print("Enter gas volume in liters: ")
	v := readInput()
	fmt.Print("Enter gas temperature in Celcius: ")
	t := readInput()

	n := (p * v) / (R * (t + 273.15))

	fmt.Println("The amount of gas in moles:", math.Floor(n * 100) / 100)

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