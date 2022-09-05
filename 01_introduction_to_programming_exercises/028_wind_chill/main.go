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

	fmt.Print("Enter air temperatures in degree Celcius: ")
	t := readInput()
	fmt.Print("Enter wind speed in km/h: ")
	v := readInput()

	wci := 13.12 + 0.6215 * t + 11.37 * math.Pow(v, 0.016) + 0.3965 * math.Pow(v, 0.016)

	fmt.Print("WCI (wind chill index): ", wci)

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