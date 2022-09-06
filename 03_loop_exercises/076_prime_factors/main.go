package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter an integer (2 or greatest): ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The prime factors of", num, "are:")
	factor := 2
	for factor <= num {
		if num % factor == 0 {
			fmt.Println(factor)
			num /= factor
		} else {
			factor += 1
		}
	}

}