package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a list of integer number with space separator: ")
	numbers := readInput()

	listOfNumbers := convertInputToIntList(numbers)

	fmt.Println("The list is sorted status:", isSorted(listOfNumbers))

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func convertInputToIntList(s string) []int {

	numbers := strings.Split(s, " ")
	var intNumbers []int
	for i := 0; i < len(numbers); i++ {
		intNum, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Fatal(err)
		}
		intNumbers = append(intNumbers, intNum)
	}

	return intNumbers

}

func isSorted(arr []int) bool {

	ascendSort := false
	descendSort := false

	ascendLoop:
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				ascendSort = true
			} else {
				ascendSort = false
				break ascendLoop
			}
		}
	}

	descendLoop:
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				descendSort = true
			} else {
				descendSort = false
				break descendLoop
			}
		}
	}

	if ascendSort || descendSort {
		return true
	} else {
		return false
	}

}