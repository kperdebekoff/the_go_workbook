package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter a first number: ")
	num1 := readNumber()
	fmt.Print("Enter a second number: ")
	num2 := readNumber()
	fmt.Print("Enter a third number: ")
	num3 := readNumber()

	median := findMediansValue(num1, num2, num3)

	fmt.Println("The median value is", median)

}

func findMediansValue(num1, num2, num3 int) int {

	if num1 > num2 {
		num2, num1 = num1, num2
	}
	if num2 > num3 {
		num3, num2 = num2, num3
	}

	return num2

}

func readNumber() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return num

}