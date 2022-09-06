package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter a numerator of a fraction: ")
	numerator := readInput()
	fmt.Print("Enter a denominator of a fraction: ")
	denominator := readInput()

	reduceNum, reduceDenom := reduceFraction(numerator, denominator)

	fmt.Println("The reduced fractioin (numerator - denominator):", reduceNum, "-", reduceDenom)

}

func reduceFraction(num, denom int) (int, int) {
 
	fraction := 0
	if num > denom {
		for i := 2; i <= denom; i++ {
			if num % i == 0 && denom % i == 0 {
				fraction = i
				break
			}
		}
	} else {
		for i := 2; i <= num; i++ {
			if num % i == 0 && denom % i == 0 {
				fraction = i
				break
			}
		}
	}

	return num / fraction, denom / fraction

}

func readInput() int {
	
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	d, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return d

}