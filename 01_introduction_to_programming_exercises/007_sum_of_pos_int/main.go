package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a positive number: ")
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	sum := (num * (num + 1)) / 2

	fmt.Println("Sum of the first n Positive integers:", sum)

}