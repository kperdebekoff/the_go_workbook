package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	const (
	centsPerToonie = 200
	centsPerLoonie = 100
	centsPerQuarter = 25
	centsPerDime = 10
	centsPerNickel = 5
	)

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("How much change to provide? (in cents): ")
	scan.Scan()
	cents, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	amountToonie := computeAmountCents(cents, centsPerToonie)
	cents = computeRemCents(cents, centsPerToonie)

	amountLoonie := computeAmountCents(cents, centsPerLoonie)
	cents = computeRemCents(cents, centsPerLoonie)
	
	amountQuarter := computeAmountCents(cents, centsPerQuarter)
	cents = computeRemCents(cents, centsPerQuarter)

	amounDime := computeAmountCents(cents, centsPerDime)
	cents = computeRemCents(cents, centsPerDime)

	amounNickel := computeAmountCents(cents, centsPerNickel)
	cents = computeRemCents(cents, centsPerNickel)

	fmt.Println("Amount of chane:")
	fmt.Println(" ", amountToonie, "two dollars coin")
	fmt.Println(" ", amountLoonie, "one dollars coin")
	fmt.Println(" ", amountQuarter, "twenty five cents coin")
	fmt.Println(" ", amounDime, "ten cents coin")
	fmt.Println(" ", amounNickel, "five cents coin")
	fmt.Println("  left", cents, "cents")

}

func computeAmountCents(cent int, centper int) int {

	return cent / centper

}

func computeRemCents(cent int, centper int) int {

	return cent % centper

}