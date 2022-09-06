package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter an integer number more than 1 (one): ")
	num := readInput()

	prime := isPrimeNumber(num)

	fmt.Println("A number prime -", prime)

}

func isPrimeNumber(n int) bool {

	count := 0
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			count += 1
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}

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