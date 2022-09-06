package main

import "fmt"

func main() {

	fmt.Println(calcTaxiFare(1.4))

}

func calcTaxiFare(distance float64) float64 {

	const (
		baseFare = 4.00
		varFare = 0.25
		traveledDist = 0.14
	)

	taxiFare := baseFare + varFare * (distance / traveledDist)

	return taxiFare

}