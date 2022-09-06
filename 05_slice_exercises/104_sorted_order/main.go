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

	fmt.Print("Enter some integer number: ")
	num := readInput()

	values := storeInSlice(num)
	sortedValues := sortSlice(values)

	fmt.Println(sortedValues)

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

func storeInSlice(num int) []int {

	var sliceOfNumbers []int

	for num != 0 {
		sliceOfNumbers = append(sliceOfNumbers, num)
		fmt.Print("Enter some integer number (0 to quit): ")
		num = readInput()
	}

	return sliceOfNumbers

}

func sortSlice(num []int) []int {

	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})

 return num

}