package main

import "fmt"

func main() {
	
	for i := 1; i <= 10000; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
		}
	}

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

func isPerfectNumber(n int) bool {

	propertyDiv := computeProperDivisor(n)
	sumDiv := 0
	for _, v := range propertyDiv {
		sumDiv += v
	}

	if sumDiv == n {
		return true
	}

	return false

}