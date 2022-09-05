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
	fmt.Print("Enter fuel efficiency for vehile in MPG (mile/gallon): ")
	scan.Scan()

	mpg, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	hlpkm := (100 * 3.78541) / (mpg * 1.60934)

	fmt.Println(mpg, "MPG (American units) is", hlpkm, "L/100 km (Canadian units)")

}