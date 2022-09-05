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
	fmt.Print("Enter some integer number: ")
	scan.Scan()
	inputNum, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}

	if inputNum % 2 != 0 {
		fmt.Println(inputNum, "is odd number")
	} else {
		fmt.Println(inputNum, "is even number")
	}
	
}