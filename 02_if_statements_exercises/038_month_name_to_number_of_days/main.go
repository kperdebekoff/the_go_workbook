package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the name of a month in lowercase: ")
	scan.Scan()
	month := scan.Text()

	switch month {

	case "january", "march", "may", "july", "august", "october", "december":
		fmt.Print("In ", month, " 31 days")
	case "february":
		fmt.Print("In ", month, " 28 or 29 days")
	case "april", "june", "september", "november":
		fmt.Print("In ", month, " 30 days")
	}

}