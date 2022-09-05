package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter a month: ")
	month := readInput()
	fmt.Print("Enter a day: ")
	day, err := strconv.Atoi(readInput())
	if err != nil {
		log.Fatal(err)
	}

	if month == "March" && day == 20 {
		fmt.Println("Season is Spring")
	} else if month == "June" && day == 21 {
		fmt.Println("Season is Summer")
	} else if month == "September" && day == 22 {
		fmt.Println("Season is Fall")
	} else if month == "December" && day == 21 {
		fmt.Println("Season is Winter")
	}

}

func readInput() string {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text()

}