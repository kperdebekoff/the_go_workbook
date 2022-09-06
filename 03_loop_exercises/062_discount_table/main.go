package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 4.95; i <= 24.95; i += 5 {
		fmt.Print("Original price $", i, "\n")
		fmt.Print("Discount amount 60%", "\n")
		fmt.Print("New price $", math.Floor(i * 0.6 * 100) / 100, "\n")
		fmt.Print("-------------------", "\n")
	}
	
}