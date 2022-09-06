package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	fmt.Print("Enter non-negative integer number: ")
	values := readInput()

	originalArray := storeDataInSlice(values)
	modifiedArray := removeExtremeValues(originalArray)

	fmt.Println("Original data:", originalArray)
	fmt.Println("Data after removed extreme values:", modifiedArray)

}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	n, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	return n

}

func storeDataInSlice(n int) []int {

	var sliceOfNumbers []int
	for n != 0 {
		sliceOfNumbers = append(sliceOfNumbers, n)
		fmt.Print("Enter non-negative integer number (0 to quit): ")
		n = readInput()
	}

	if len(sliceOfNumbers) < 4 {
		log.Fatal("Entered less than 4 non-negative integer number!")
	}

	return sliceOfNumbers

}

func removeExtremeValues(n []int) []int {

	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})

	n = n[1:len(n) - 1]

	return n

}