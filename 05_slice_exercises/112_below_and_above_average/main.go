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

	sliceOfNum := storeDataInSlice(number)
	average := calcAverage(sliceOfNum)
	belowAverage, equalAverage, aboveAverage := changeDataOrder(sliceOfNum, average)

	fmt.Println("List of entered numbers:", sliceOfNum)
	fmt.Println("Average of entered numbers:", average)
	fmt.Println("Numbers below average:", belowAverage)
	fmt.Println("Numbers equal average:", equalAverage)
	fmt.Println("Numbers above average:", aboveAverage)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func storeDataInSlice(s string) []int {

	var numbers []int

	for s != "" {

		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)

		fmt.Print("Enter a positive integer number (blank to quit): ")
		s = readInput()

	}

	return numbers

}

func calcAverage(n []int) int {

	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum / len(n)

}

func changeDataOrder(numbers []int, average int) ([]int, []int, []int) {

	var belowAverage []int
	var aboveAverage []int
	var equalAverage []int

	for _, v := range numbers {
		if v < average {
			belowAverage = append(belowAverage, v)
		} else if v > average {
			aboveAverage = append(aboveAverage, v)
		} else {
			equalAverage = append(equalAverage, v)
		}
	}

	return belowAverage, equalAverage, aboveAverage

}