package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter a postal code: ")
	postalCode := readInput()

	fmt.Println(displayProvice(postalCode))

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}

func displayProvice(postalCode string) string {

	arrRuneCode := []rune(postalCode)

	address := ""
	charProvince := map[string]string{
		"A": "Newfoundland",
		"B": "Nova Scotia",
		"C": "Price Edward Island",
		"E": "New Brunswick",
		"G": "Quebec",
		"H": "Quebec",
		"J": "Quebec",
		"K": "Ontario",
		"L": "Ontario",
		"M": "Ontario",
		"N": "Ontario",
		"P": "Ontario",
		"R": "Manitoba",
		"S": "Saskatchewan",
		"T": "Alberta",
		"V": "British Columbia",
		"X": "Nunavut or Nortwest Territories",
		"Y": "Yukon",
	}

	isRightNum := func(arr []rune) bool {

		isRight := false
		if len(arr) == 6 {
			for i := 1; i < len(arr); i += 2 {
				if arr[i] >= 48 && arr[i] <= 57 {
					isRight = true
				} else {
					isRight = false
					break
				}
			}
		}
		
		return isRight

	}

	isRightSymb := func(arr []rune) bool {

		wrongSymb := []rune{68, 70, 73, 79, 81, 87, 90}
		isRight := true
		if len(arr) == 6 {
			Loop:
			for i := 0; i < len(arr); i += 2 {
				if arr[i] >= 60 && arr[i] <= 90 {
					for j := 0; j < len(wrongSymb); j++ {
						if arr[i] == wrongSymb[j] {
							isRight = false
							break Loop
						}
					}
				} else {
					isRight = false
				}
			}
		} else {
			isRight = false
		}

		return isRight

	}
	
	if isRightNum(arrRuneCode) && isRightSymb(arrRuneCode) {
		if arrRuneCode[1] > 48 {
			address += "The address is urban in "
		} else {
			address += "The address is rural in "
		}

		for k, v := range charProvince {
			if k == string(arrRuneCode[0]) {
				address += v
			}
		}
	} else {
		address = "Possible, your postal code begins with an invalid charachter"
		address += " or has invalid length! Please enter a valid postal code!"
	}

	return address

}