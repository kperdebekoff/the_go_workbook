package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func inputData() float64 {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	data, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {

	fmt.Print("Enter a total number of seconds: ")
	s := int(inputData())

	days := s / (24 * 3600)
	s = s % (24 * 3600)
	hours := s / 3600
	s = s % 3600
	minutes := s / 60
	s = s % 60

	fmt.Print(days, "d:", hours, "h:", minutes, "m:", s, "s")
	
}