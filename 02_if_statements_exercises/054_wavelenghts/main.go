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
	fmt.Print("Enter a wavelength of the visible spectrum: ")
	scan.Scan()
	wavelength, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	if wavelength >= 380 && wavelength < 450 {
		fmt.Println("Violet color spectrum")
	} else if wavelength >= 450 && wavelength < 495 {
		fmt.Println("Blue color spectrum")
	} else if wavelength >= 495 && wavelength < 570 {
		fmt.Println("Green color spectrum")
	} else if wavelength >= 570 && wavelength < 590 {
		fmt.Println("Yellow color spectrum")
	} else if wavelength >= 590 && wavelength < 620 {
		fmt.Println("Orange color spectrum")
	} else if wavelength >= 620 && wavelength <= 750 {
		fmt.Println("Red color spectrum")
	} else {
		fmt.Println("Entered wavelength outside of the visible spectrum!")
	}

}