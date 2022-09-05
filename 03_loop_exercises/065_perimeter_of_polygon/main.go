package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	perimeter := 0.0
	fmt.Print("Enter the x part of the coordinate: ")
	xFirst := readInput()
	fmt.Print("Enter the y part of the coordinate: ")
	yFirst := readInput()

	xa := xFirst
	ya := yFirst
	for {
		fmt.Print("Enter the x part of the coordinate (blank to quit): ")
		xb := readInput()
		if xb == -1 {
			break
		}
		fmt.Print("Enter the y part of the coordinate: ")
		yb := readInput()

		dist := math.Sqrt(math.Pow((xb - xa), 2) + math.Pow((yb - ya), 2))
		perimeter += dist

		xa = xb
		ya = yb
	}

	dist := math.Sqrt(math.Pow((xFirst - xa), 2) + math.Pow((yFirst - ya), 2))
	perimeter += dist

	fmt.Println("The perimeter of that polygon:", perimeter)
}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	part, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		return -1
	}
	return part

}