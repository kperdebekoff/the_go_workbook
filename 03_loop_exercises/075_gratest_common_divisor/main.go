package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter a first number: ")
	num1 := readInput()
	fmt.Print("Enter a second number: ")
	num2 := readInput()

	d := 0
	if num1 < num2 {
		d = num1
	} else {
		d = num2
	}

	for num1 % d != 0  || num2 % d != 0 {
		d -= 1
	}

	fmt.Println("Greatest common divisor of", num1, "and", num2, ":", d)
}

func readInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return num

}