package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter a positive integer number: ")
	number := readInput()

	properDiv := computeProperDivisor(number)

	fmt.Println(properDiv)
}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	return num

}

func computeProperDivisor(num int) []int {

	var divisors []int
	for i := 1; i < num; i++ {
		if num % i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
	
}