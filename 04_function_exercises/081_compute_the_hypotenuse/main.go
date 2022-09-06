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
	
	fmt.Print("Enter fist shorter side of a right triangle: ")
	side1 := readTriangleSide()
	fmt.Print("Enter second shorter side of a right triangle: ")
	side2 := readTriangleSide()

	hypotenuse := calcHypotenuse(side1, side2)

	fmt.Println("Length of the hypotenuse:", hypotenuse)

}

func calcHypotenuse(side1, side2 float64) float64 {

	return math.Sqrt(side1 * side1 + side2 * side2)

}

func readTriangleSide() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	side, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return side

}