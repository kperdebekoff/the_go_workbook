package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a month and day (for example: March 21): ")
	scan.Scan()
	holiday := scan.Text()

	switch holiday {

	case "January 1":
		fmt.Print("New year's day")
	case "July 1":
		fmt.Print("Canada day")
	case "December 25":
		fmt.Print("Christmas day")
	default:
		fmt.Print("Entered month and day don't correspond to a fixed-date holiday")
		
	}

}