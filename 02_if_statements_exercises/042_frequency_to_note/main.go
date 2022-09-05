package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a note frequency: ")
	scan.Scan()
	freq, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	if 260.63 <= freq && freq <= 262.63 {
		fmt.Println("C4")
	} else if 292.66 <= freq && freq <= 294.66 {
		fmt.Println("D4")
	} else if 328.63 <= freq && freq <= 330.63 {
		fmt.Println("E4")
	} else if 348.23 <= freq && freq <= 350.23 {
		fmt.Println("F4")
	} else if 391.00 <= freq && freq <= 393.00 {
		fmt.Println("G4")
	} else if 439.00 <= freq && freq <= 441.00 {
		fmt.Println("A4")
	} else if 492.88 <= freq && freq <= 494.88 {
		fmt.Println("B4")
	} else {
		fmt.Println("Frequensy doesn't correspond to a known note")
	}
	
}