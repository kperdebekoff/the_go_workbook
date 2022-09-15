package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter an integer number: ")
	num := readInput()

	numbers := storeDataInSlice(num)
	orderedNumbers := changeDataOrder(numbers)

	for _, v := range orderedNumbers {
		fmt.Println(v)
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func storeDataInSlice(n string) []int {

	var nums []int
	for n != "" {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)

		fmt.Print("Enter an integer number (blank to quit): ")
		n = readInput()
	}

	return nums

}

func changeDataOrder(nums []int) []int {

	var orderedNumbers []int

	for _, v := range nums {
		if v < 0 {
			orderedNumbers = append(orderedNumbers, v)
		}
	}

	for _, v := range nums {
		if v == 0 {
			orderedNumbers = append(orderedNumbers, v)
		}
	}
	
	for _, v := range nums {
		if v > 0 {
			orderedNumbers = append(orderedNumbers, v)
		}
	}

	return orderedNumbers

}