package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter a decimal number to convert other base: ")
	num1 := readInput()
	fmt.Print("Enter convertion base of number: ")
	base := readInput()

	fmt.Print("Enter any number to convert decimal: ")
	num2 := readInput()
	fmt.Print("Enter base of entered number: ")
	base2 := readInput()

	n1 := convertDecToOtherBase(base, num1)
	n2 := convertOtherBaseToDec(base2, num2)

	fmt.Println("----------------------------------------------------")
	fmt.Println("Decimal number", num1, "to base", base, "is", n1)
	fmt.Println("Number", num2, "with base", base, "to decimal is ", n2)

}

func convertDecToOtherBase(b, n string) string {

	strNum := ""

	bases, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}

	rem := 0
	for num != 0 {
		rem = num % bases
		switch rem {
		case 10:
			strNum += "A"
		case 11:
			strNum += "B"
		case 12:
			strNum += "C"
		case 13:
			strNum += "D"
		case 14:
			strNum += "E"
		case 15:
			strNum += "F"
		default:
			strNum += strconv.Itoa(rem)
		}
		num /= bases
	}
	revNum := ""
	for i := 0; i < len(strNum); i++ {
		revNum += string(strNum[len(strNum) - i - 1])
	}
	return revNum

}

func convertOtherBaseToDec(b, n string) int {

	dec := 0.0

	bases, err := strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(n); i++ {
		num, err := strconv.ParseFloat(string(n[i]), 64)
		if err != nil {
			switch string(n[i]) {
			case "A", "a":
				dec += 10 * math.Pow(bases, float64(len(n) - 1 - i))
			case "B", "b":
				dec += 11 * math.Pow(bases, float64(len(n) - 1 - i))
			case "C", "c":
				dec += 12 * math.Pow(bases, float64(len(n) - 1 - i))
			case "D", "d":
				dec += 13 * math.Pow(bases, float64(len(n) - 1 - i))
			case "E", "e":
				dec += 14 * math.Pow(bases, float64(len(n) - 1 - i))
			case "F", "f":
				dec += 15 * math.Pow(bases, float64(len(n) - 1 - i))
			default:
				log.Fatal(err)
			}
		}
		dec += num * math.Pow(bases, float64(len(n) - 1 - i))
	}
	return int(dec)

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}