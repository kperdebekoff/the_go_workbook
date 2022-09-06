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
	
	fmt.Print("Enter a hexadecimal number for convert to decimal: ")
	hexN := readInput()
	fmt.Print("Enter a decimal number for convert to hexadecimal: ")
	decN := readInput()

	hexToDec := convertHexToInt(hexN)
	decToHex := convertIntToHex(decN)
	
	fmt.Println(hexToDec)
	fmt.Println(decToHex)

}


func convertHexToInt(n string) int {

	dec := 0.0

	for i := 0; i < len(n); i++ {
		num, err := strconv.ParseFloat(string(n[i]), 64)
		if err != nil {
			switch string(n[i]) {
			case "A", "a":
				dec += 10 * math.Pow(16, float64(len(n) - 1 - i))
			case "B", "b":
				dec += 11 * math.Pow(16, float64(len(n) - 1 - i))
			case "C", "c":
				dec += 12 * math.Pow(16, float64(len(n) - 1 - i))
			case "D", "d":
				dec += 13 * math.Pow(16, float64(len(n) - 1 - i))
			case "E", "e":
				dec += 14 * math.Pow(16, float64(len(n) - 1 - i))
			case "F", "f":
				dec += 15 * math.Pow(16, float64(len(n) - 1 - i))
			default:
				log.Fatal(err)
			}
		}
		dec += num * math.Pow(16, float64(len(n) - 1 - i))
	}
	return int(dec)

}

func convertIntToHex(n string) string {

	num, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}
	hex := ""
	r := 0
	for num != 0 {
		r = num % 16
		switch r {
		case 10:
			hex += "A"
		case 11:
			hex += "B"
		case 12:
			hex += "C"
		case 13:
			hex += "D"
		case 14:
			hex += "E"
		case 15:
			hex += "F"
		default:
			hex += strconv.Itoa(r)
		}
		num /= 16
	}

	h := ""
	for i := 0; i < len(hex); i++ {
		h += string(hex[len(hex) - i - 1])
	}

	return h

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}