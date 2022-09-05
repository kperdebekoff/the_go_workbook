package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main () {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a radius of a circle in meter: ")
	scan.Scan()
	r, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	
	area := math.Pi * math.Pow(r, 2)
	volume := (4 / 3) * math.Pi *  math.Pow(r, 3)

	fmt.Println("The area of a circle is:", area, "m^2")
	fmt.Println("The volume of a circle is:", volume, "m^3")
	
}