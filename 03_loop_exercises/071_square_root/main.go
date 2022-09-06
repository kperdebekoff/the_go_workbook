package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fmt. Print("Enter a number for compute the square root: ")
	x := readInput()
	guess := x / 2

	for {
		guess = (guess + x / guess) / 2
		if guess * guess == x {
			break
		}
	}

	fmt.Println("The square root of", x, "equal", guess)

}

func readInput() float64 {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return num

}