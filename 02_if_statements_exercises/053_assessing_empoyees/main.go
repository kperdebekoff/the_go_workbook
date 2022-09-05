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
	fmt.Print("Enter your rating: ")
	scan.Scan()
	rating, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}

	if rating >= 0.6 {
		raise := rating * 2400
		fmt.Println("Meritorius perfomance - the amount of raise $", raise)
	} else if rating == 0.4 {
		raise := rating * 2400
		fmt.Println("Acceptable perfomance - the amount of raise $", raise)
	} else if rating == 0.0 {
		raise := rating * 2400
		fmt.Println("Unacceptable perfomance - the amount of raise $", raise)
	} else {
		fmt.Println("Entered invalid rating!")
	}

}