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

	v0 := 0
	g := 9.8

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the height in meter: ")
	scan.Scan()
	height, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	vf := math.Sqrt(math.Pow(float64(v0), 2) + 2 * g * height)

	fmt.Println("Final speed:", vf, "m/s")
	
}