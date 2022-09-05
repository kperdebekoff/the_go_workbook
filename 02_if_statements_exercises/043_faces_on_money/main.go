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
	fmt.Print("Enter denomination of a USA banknote: ")
	scan.Scan()
	banknote, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	switch banknote {

	case 1:
		fmt.Print("$", banknote, " - George Washington")
	case 2:
		fmt.Print("$", banknote, " - Thomas Jefferson")
	case 5:
		fmt.Print("$", banknote, " - Abraham Lincoln")
	case 10:
		fmt.Print("$", banknote, " - Alexander Hamilton")
	case 20:
		fmt.Print("$", banknote, " - Andrew Jackson")
	case 50:
		fmt.Print("$", banknote, " - Ulysses S. Grant")
	case 100:
		fmt.Print("$", banknote, " - Benjamin Franklin")
	default:
		fmt.Println("No such banknote exists")
	}

}