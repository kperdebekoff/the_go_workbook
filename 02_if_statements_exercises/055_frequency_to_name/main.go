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
	fmt.Print("Enter a frequency of the radiation in Hz: ")
	scan.Scan()
	freq, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	if freq < 3 * math.Pow(10, 9) {
		fmt.Println("Radio waves")
	} else if freq >= 3 * math.Pow(10, 9) && freq < 3 * math.Pow(10, 12) {
		fmt.Println("Microwaves")
	} else if freq >= 3 * math.Pow(10, 12) && freq < 4.3 * math.Pow(10, 14) {
		fmt.Println("Infrared light")
	} else if freq >= 4.3 * math.Pow10(14) && freq < 7.5 * math.Pow10(14) {
		fmt.Println("Visible light")
	} else if freq >= 7.5 * math.Pow10(14) && freq < 3 * math.Pow10(17) {
		fmt.Println("Ultraviolet light")
	} else if freq >= 3 * math.Pow10(17) && freq < 3 * math.Pow10(19) {
		fmt.Println("X-rays")
	} else if freq >= 3 * math.Pow10(19) {
		fmt.Println("Gamma rays")
	}

}