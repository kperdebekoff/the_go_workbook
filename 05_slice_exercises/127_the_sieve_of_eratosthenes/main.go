package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt.Print("Enter a limit of number to find all of the prime numbers: ")
	limit := readInput()

	primeNumbers := eratostheneSieve(limit)
	for _, v := range primeNumbers {
		if v != 0 {
			fmt.Println(v)
		}
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

func eratostheneSieve(limit int) []int {

	numbers := []int{}
	for i := 0; i <= limit; i++ {
		numbers = append(numbers, i)
	}
	numbers[1] = 0

	primeNums := func(numbers []int, p int) []int {

		for i := 0; i < len(numbers); i += 1 {
			if numbers[i] % p == 0 && numbers[i] != p {
				numbers[i] = 0
			}
		}
		return numbers

	}
	p := 2
	primeNums(numbers, p)

	for i := p; i < len(numbers); i++ {
		if numbers[i] % p != 0 && numbers[i] != p {
			p = numbers[i]
		}
		primeNums(numbers, p)
	}

	return numbers

}