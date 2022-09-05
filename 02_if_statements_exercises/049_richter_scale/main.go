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
	fmt.Print("Enter a magnitude ranges on the Richter scale: ")
	scan.Scan()
	magnitude, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	if magnitude < 2.0 {
		fmt.Println("Micro")
	} else if magnitude >= 2.0 && magnitude < 3.0 {
		fmt.Println("Very minor")
	} else if magnitude >= 3.0 && magnitude < 4.0 {
		fmt.Println("Minor")
	} else if magnitude >= 4.0 && magnitude < 5.0 {
		fmt.Println("Light")
	} else if magnitude >= 5.0 && magnitude < 6.0 {
		fmt.Println("Moderate")
	} else if magnitude >= 6.0 && magnitude < 7.0 {
		fmt.Println("Strong")
	} else if magnitude >= 7.0 && magnitude < 8.0 {
		fmt.Println("Major")
	} else if magnitude >= 8.0 && magnitude < 10.0 {
		fmt.Println("Great")
	} else if magnitude >= 10.0 {
		fmt.Println("Meteoric")
	}

}