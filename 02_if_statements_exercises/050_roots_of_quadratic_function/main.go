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

	fmt.Print("Enter the first 'a' value of quadratic function: ")
	a := readInput()
	fmt.Print("Enter the second 'b' value of quadratic function: ")
	b := readInput()
	fmt.Print("Enter the third 'c' value of quadratic function: ")
	c := readInput()

	d := math.Pow(b, 2) - 4 * a * c

	if d < 0 {
		fmt.Println("The quadratic equation does not have any real roots")
	} else if d == 0 {
		root := -b / 2 * a
		fmt.Println("The quadratic equation have one root", root)
	} else {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("First root:", root1)
		fmt.Println("Second root:", root2)
	}

}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	value, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	return value

}
