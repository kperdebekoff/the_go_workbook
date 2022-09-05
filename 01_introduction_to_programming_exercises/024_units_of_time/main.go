package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a number of days: ")
	d := readInput()
	fmt.Print("Enter a number of hours: ")
	h := readInput()
	fmt.Print("Enter a number of minutes: ")
	m := readInput()
	fmt.Print("Enter a number of seconds: ")
	s := readInput()

	amountSeconds := (d * 24 * 3600) + (h * 3600) + (m * 60) + s

	fmt.Print("Total number of seconds: ", amountSeconds)

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