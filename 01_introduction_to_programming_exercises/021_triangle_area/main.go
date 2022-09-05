package main

import (
	"bufio"
	"os"
	"strconv"
	"log"
	"fmt"
)

func main() {

	fmt.Print("Enter a trinagle length of the base: ")
	length := readInput()
	fmt.Print("Enter a trinagle height: ")
	height := readInput()

	area := (length * height) / 2

	fmt.Println("Triangle area: ", area)

}

func readInput() float64 {
	
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	data, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return data

}