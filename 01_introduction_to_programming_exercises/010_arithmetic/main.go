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
	fmt.Print("Enter the first number: ")
	a := convertStrToInt(readInput())
	fmt.Print("Enter the second number: ")
	b := convertStrToInt(readInput())

	sum := a + b
	diff := a - b
	mult := a * b
	div := float64(a / b)
	rem := float64(a % b)
	lgA := math.Log10(float64(a))
	powA := math.Pow(float64(a), float64(b))

	fmt.Println("Sum a and b:", sum)
	fmt.Println("Difference a and b:", diff)
	fmt.Println("The product a and b:", mult)
	fmt.Println("The quotient a divide b:", div)
	fmt.Println("The remainder a divide b:", rem)
	fmt.Println("The log10 from a:", lgA)
	fmt.Println("The result a power b:", powA)

}

func convertStrToInt(a string) int {

	num, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return num

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}