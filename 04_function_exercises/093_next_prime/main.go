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

	nxtPrmNm := nextPrimeNum(num)	

	fmt.Println("Next prime number is -", nxtPrmNm)

}

func nextPrimeNum(n int) int {

	for {
		n += 1
		prime := isPrimeNumber(n)
		if prime {
			return n
		}
	}

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