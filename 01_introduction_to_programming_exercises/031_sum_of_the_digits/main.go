package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a four-digit integer number: ")
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	last := num % 10
	preLast := (num / 10) % 10
	second := (num / 100) % 10
	first := num / 1000
	sumDig := last + preLast + second + first

	fmt.Println("Sum of the digits in the number:", sumDig)
	
}