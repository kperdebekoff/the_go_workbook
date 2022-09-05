package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter the first side length of a triangle: ")
	firstSide := readInput()
	fmt.Print("Enter the second side length of a triangle: ")
	secondSide := readInput()
	fmt.Print("Enter the third side length of a triangle: ")
	thirdSide := readInput()

	if firstSide == secondSide && secondSide == thirdSide {
		fmt.Println("Equlateral triangle")
	} else if firstSide != secondSide && secondSide != thirdSide && firstSide != thirdSide {
		fmt.Println("Scalene triangle")
	} else {
		fmt.Println("Isosceles triangle")
	}

}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	side, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	return side

}