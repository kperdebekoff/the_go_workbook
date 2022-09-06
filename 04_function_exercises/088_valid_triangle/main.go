package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	fmt.Print("Enter a first segment length: ")
	s1 := readUserInput()
	fmt.Print("Enter a second segment length: ")
	s2 := readUserInput()
	fmt.Print("Enter a third segment length: ")
	s3 := readUserInput()

	isTriangle := verifyTriangle(s1, s2, s3)

	fmt.Println("Is it a valid triangle? -", isTriangle)

}

func verifyTriangle(s1, s2, s3 int) bool {

	isTriangle := false
	if s1 <= 0 || s2 <= 0 || s3 <= 0 {
		isTriangle = false
	} else {
		if s1 <= s2 + s3 || s2 <= s1 + s3 || s3 <= s1 + s2 {
			isTriangle = true
		} else {
			isTriangle = false
		}
	}
	return isTriangle

}

func readUserInput() int {

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	segment, err := strconv.Atoi(scan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return segment

}