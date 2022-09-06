package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	fmt.Print("Enter a binary number: ")
	binaryNum := readInput()

	var decimal float64 = 0
	for i := 0; i < len(binaryNum); i++ {
		decimal += (float64(binaryNum[i] - '0') * math.Pow(2, float64(len(binaryNum)) - float64(i) - 1))
	}

	fmt.Print(decimal)

}

func readInput() []rune {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return []rune(scan.Text())

}